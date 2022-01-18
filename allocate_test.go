package bitswap

import (
	"fmt"
	standard "github.com/bdengine/go-ipfs-blockchain-standard/model"
	blocks "github.com/ipfs/go-block-format"
	"testing"
)

func TestAllocateBlocks(t *testing.T) {
	var blockList []blocks.Block
	for i := 0; i < 3; i++ {
		blockList = append(blockList, blocks.NewBlock([]byte{byte(i)}))
	}
	var serverList []standard.CorePeer
	for i := 0; i < 3; i++ {
		orgName := fmt.Sprintf("org%v", i)
		for j := 0; j < 1; j++ {
			serverList = append(serverList, standard.CorePeer{
				Org: orgName,
				Peer: standard.Peer{
					PeerId: fmt.Sprintf("%speer%v", orgName, j),
				},
			})
		}
	}
	peerInfo, blockTarList, orgInfo, err := findAllocateCondition(blockList, serverList)
	if err != nil {
		t.Fatal(err)
	}
	res, err := allocateBlocksAlgorithm(peerInfo, blockTarList, orgInfo)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	for s, message := range res {
		fmt.Printf("节点：%v，获得分块", s)
		b := message.Blocks()
		for _, block := range b {
			fmt.Printf("%v(cid:%s),", block.RawData()[0], block.Cid().String())
		}
		fmt.Printf("\n")
	}
}
