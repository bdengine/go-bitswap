package message

import (
	pb "github.com/ipfs/go-bitswap/message/pb"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type uuid string
type idHash string

type implPlus struct {
	full           bool
	wantlist       map[cid.Cid]*Entry
	blocks         map[cid.Cid]blocks.Block
	blockPresences map[cid.Cid]pb.Message_BlockPresenceType
	pendingBytes   int32

	// 备份信息
	backupblocks map[cid.Cid]backupLoad
	//
	deleteblocks map[uuid][]cid.Cid
}

type backupLoad struct {
	// 目标节点
	targetPeer []peer.ID
	// 文件唯一id和文件片cid的hash值
	idHash string
	// 文件片
	block blocks.Block
}

type deleteLoad struct {
	// 文件唯一id
	uuid string
	// 文件片cid
	cid cid.Cid
}

type storeLoad struct {
	idHashMap  map[idHash]uuid
	storePeer  []peer.ID
	deleteFlag bool
}

func StoreBlock() {
	// 读取目标节点是否包本节点
	// 包含，levelDb存储相关信息
	//
}

func DeleteBlocks(id uuid, cidList []cid.Cid) {
	for _, c := range cidList {
		deleteBlock(c, id)
	}
}

func deleteBlock(cid cid.Cid, id uuid) {
	// levelDb 查询cid的备份信息
	// 无相关信息返回无
	// id+cid hash计算idHash
	// 检查ownerMap中是否存在该idHash
	// 存在,删除该idHash
	// 删除后，若idHash为空，直接删除对应block
	// 操作完成过后，向外分发删除信息
}
