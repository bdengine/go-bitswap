package bitswap

import (
	"context"
	"fmt"
	engine "github.com/bdengine/go-bitswap/internal/decision"
	bsmsg "github.com/bdengine/go-bitswap/message"
	pb "github.com/bdengine/go-bitswap/message/pb"
	cid "github.com/ipfs/go-cid"
	process "github.com/jbenet/goprocess"
	procctx "github.com/jbenet/goprocess/context"
	"github.com/libp2p/go-libp2p-core/peer"
	"go.uber.org/zap"
)

// TaskWorkerCount is the total number of simultaneous threads sending
// outgoing messages
var TaskWorkerCount = 8

func (bs *Bitswap) startWorkers(ctx context.Context, px process.Process) {

	// Start up workers to handle requests from other nodes for the data on this node
	for i := 0; i < TaskWorkerCount; i++ {
		i := i
		px.Go(func(px process.Process) {
			bs.taskWorker(ctx, i)
		})
	}

	if bs.provideEnabled {
		// Start up a worker to manage sending out provides messages
		px.Go(func(px process.Process) {
			bs.provideCollector(ctx)
		})

		// Spawn up multiple workers to handle incoming blocks
		// consider increasing number if providing blocks bottlenecks
		// file transfers
		px.Go(bs.provideWorker)
	}
}

func (bs *Bitswap) taskWorker(ctx context.Context, id int) {
	defer log.Debug("bitswap task worker shutting down...")
	log := log.With("ID", id)
	for {
		log.Debug("Bitswap.TaskWorker.Loop")
		select {
		case nextEnvelope := <-bs.engine.Outbox():
			select {
			case envelope, ok := <-nextEnvelope:
				if !ok {
					continue
				}
				err := bs.sendBlocks(ctx, envelope)
				if err == nil {
					// 无错误时记录发送消息
					bs.engine.MessageSent(envelope.Peer, envelope.Message)
					if bs.wiretap != nil {
						bs.wiretap.MessageSent(envelope.Peer, envelope.Message)
					}
				} else {
					// 记录错误和重试
				}
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (bs *Bitswap) logOutgoingBlocks(env *engine.Envelope) {
	if ce := sflog.Check(zap.DebugLevel, "sent message"); ce == nil {
		return
	}

	self := bs.network.Self()

	for _, blockPresence := range env.Message.BlockPresences() {
		c := blockPresence.Cid
		switch blockPresence.Type {
		case pb.Message_Have:
			log.Debugw("sent message",
				"type", "HAVE",
				"cid", c,
				"local", self,
				"to", env.Peer,
			)
		case pb.Message_DontHave:
			log.Debugw("sent message",
				"type", "DONT_HAVE",
				"cid", c,
				"local", self,
				"to", env.Peer,
			)
		default:
			panic(fmt.Sprintf("unrecognized BlockPresence type %v", blockPresence.Type))
		}

	}
	for _, block := range env.Message.Blocks() {
		log.Debugw("sent message",
			"type", "BLOCK",
			"cid", block.Cid(),
			"local", self,
			"to", env.Peer,
		)
	}
}

/* 描述: 向指定的节点发送消息
 * 前置状态: message内容正确填装，节点正确且能连接到
 * 后置状态: 发送message，忽略是否成功
 */
func (bs *Bitswap) SendBlocks(ctx context.Context, idString string, message bsmsg.BitSwapMessage) {
	id, _ := peer.Decode(idString)
	if !message.Empty() {
		envelope := &engine.Envelope{
			Peer:    id,
			Message: message,
			// 发送成功的回调，暂时设置为空
			Sent: func() {
				return
			},
		}
		bs.sendBlocks(ctx, envelope)
	}
}

func (bs *Bitswap) sendBlocks(ctx context.Context, env *engine.Envelope) error {
	// Blocks need to be sent synchronously to maintain proper backpressure
	// throughout the network stack
	defer env.Sent()

	err := bs.network.SendMessage(ctx, env.Peer, env.Message)
	if err != nil {
		log.Debugw("failed to send blocks message",
			"peer", env.Peer,
			"error", err,
		)
		return err
	}

	bs.logOutgoingBlocks(env)

	dataSent := 0
	blocks := env.Message.Blocks()
	for _, b := range blocks {
		dataSent += len(b.RawData())
	}
	bs.counterLk.Lock()
	bs.counters.blocksSent += uint64(len(blocks))
	bs.counters.dataSent += uint64(dataSent)
	bs.counterLk.Unlock()
	bs.sentHistogram.Observe(float64(env.Message.Size()))
	log.Debugw("sent message", "peer", env.Peer)
	return nil
}

func (bs *Bitswap) provideWorker(px process.Process) {
	// FIXME: OnClosingContext returns a _custom_ context type.
	// Unfortunately, deriving a new cancelable context from this custom
	// type fires off a goroutine. To work around this, we create a single
	// cancelable context up-front and derive all sub-contexts from that.
	//
	// See: https://github.com/ipfs/go-ipfs/issues/5810
	ctx := procctx.OnClosingContext(px)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	limit := make(chan struct{}, provideWorkerMax)

	limitedGoProvide := func(k cid.Cid, wid int) {
		defer func() {
			// replace token when done
			<-limit
		}()

		log.Debugw("Bitswap.ProvideWorker.Start", "ID", wid, "cid", k)
		defer log.Debugw("Bitswap.ProvideWorker.End", "ID", wid, "cid", k)

		ctx, cancel := context.WithTimeout(ctx, provideTimeout) // timeout ctx
		defer cancel()

		if err := bs.network.Provide(ctx, k); err != nil {
			log.Warn(err)
		}
	}

	// worker spawner, reads from bs.provideKeys until it closes, spawning a
	// _ratelimited_ number of workers to handle each key.
	for wid := 2; ; wid++ {
		log.Debug("Bitswap.ProvideWorker.Loop")

		select {
		case <-px.Closing():
			return
		case k, ok := <-bs.provideKeys:
			if !ok {
				log.Debug("provideKeys channel closed")
				return
			}
			select {
			case <-px.Closing():
				return
			case limit <- struct{}{}:
				go limitedGoProvide(k, wid)
			}
		}
	}
}

func (bs *Bitswap) provideCollector(ctx context.Context) {
	defer close(bs.provideKeys)
	var toProvide []cid.Cid
	var nextKey cid.Cid
	var keysOut chan cid.Cid

	for {
		select {
		case blkey, ok := <-bs.newBlocks:
			if !ok {
				log.Debug("newBlocks channel closed")
				return
			}

			if keysOut == nil {
				nextKey = blkey
				keysOut = bs.provideKeys
			} else {
				toProvide = append(toProvide, blkey)
			}
		case keysOut <- nextKey:
			if len(toProvide) > 0 {
				nextKey = toProvide[0]
				toProvide = toProvide[1:]
			} else {
				keysOut = nil
			}
		case <-ctx.Done():
			return
		}
	}
}
