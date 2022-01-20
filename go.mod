module github.com/bdengine/go-bitswap

require (
	github.com/Hyperledger-TWGC/tjfoc-gm v1.4.0
	github.com/bdengine/go-ipfs-blockchain-selector v0.0.4
	github.com/bdengine/go-ipfs-blockchain-standard v0.0.2
	github.com/cskr/pubsub v1.0.2
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-detect-race v0.0.1
	github.com/ipfs/go-ipfs-blockstore v0.1.4
	github.com/ipfs/go-ipfs-blocksutil v0.0.1
	github.com/ipfs/go-ipfs-delay v0.0.1
	github.com/ipfs/go-ipfs-exchange-interface v0.0.1
	github.com/ipfs/go-ipfs-routing v0.1.0
	github.com/ipfs/go-ipfs-util v0.0.2
	github.com/ipfs/go-log v1.0.4
	github.com/ipfs/go-metrics-interface v0.0.1
	github.com/ipfs/go-peertaskqueue v0.2.0
	github.com/jbenet/goprocess v0.1.4
	github.com/libp2p/go-buffer-pool v0.0.2
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-loggables v0.1.0
	github.com/libp2p/go-libp2p-netutil v0.1.0
	github.com/libp2p/go-libp2p-testing v0.4.0
	github.com/libp2p/go-msgio v0.0.6
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multistream v0.2.0
	go.uber.org/zap v1.16.0
	google.golang.org/protobuf v1.23.0
)

replace (
	github.com/ipfs/go-cid => github.com/bdengine/go-cid latest
	github.com/ipfs/go-peertaskqueue => github.com/bdengine/go-peertaskqueue latest
	github.com/bdengine/go-ipfs-blockchain-selector => github.com/bdengine/go-ipfs-blockchain-selector latest
	github.com/bdengine/go-ipfs-blockchain-standard => github.com/bdengine/go-ipfs-blockchain-standard latest
)

go 1.15
