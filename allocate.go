package bitswap

import (
	"context"
	"fmt"
	bsmsg "github.com/ipfs/go-bitswap/message"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs-auth/standard"
)

const (
	DefaultBackupTar = 2
)

// 返回结果 []backupResult
type BackupResult struct {
	OrgName string
	PeerId  string
	Blocks  Backup
}

// 约束结构
// len(backup) < max - 1
// backup[any] < peerNum
type OrgBackupInfo struct {
	Backup  Backup
	OrgName string
	PeerNum int
}

type Backup map[blocks.Block]int

type BlockBackupTar struct {
	Block     blocks.Block
	BackupTar int
}

func BackupSum(b1 Backup, b2 Backup) Backup {
	for block, i := range b2 {
		b1[block] += i
	}
	return b1
}

/* 描述: 计算文件块落点并包装成消息
 * 入参:[文件各节点备份数量，各文件块需备份数量，各组织文件备份数量]
 * 出参：[发往各节点的消息]
 */
// allocateBlocksAlgorithm serverList:各节点备份情况，blockList：分片列表，backupTar:备份目标数量
func allocateBlocksAlgorithm(serverList []BackupResult, bList []BlockBackupTar, orgInfo map[string]OrgBackupInfo) (map[string]bsmsg.BitSwapMessage, error) {

	blockNum := len(bList)

	serverNum := len(serverList)
	// todo 检查组织数和备份数是否满足条件

	res := map[string]bsmsg.BitSwapMessage{}
	// 开始循环分配
	for i1, tar := range bList {
		// 错开查询的起始
		i2 := i1
		// 有效循环次数
		times := 0
		// 当前备份数量
		backupNum := 0
		// 最多循环节点数的次数，保证每一次都查询了所有节点
		for ; times < serverNum && backupNum < tar.BackupTar; times++ {
			i2 %= serverNum
			server := serverList[i2]
			orgName := server.OrgName
			info := orgInfo[orgName]
			// 该组织已有该文件片备份数大于等于节点数，不再分配
			if info.Backup[tar.Block] > info.PeerNum {
				i2++
				continue
			}
			// 该组织已拥有max-1个分片，不再分配
			if len(info.Backup) >= blockNum-1 {
				i2++
				continue
			}
			// 该节点已经拥有该文件片，不再分配

			if server.Blocks[tar.Block] > 0 {
				i2++
				continue
			}
			// 将该块分配给该节点，组织记录
			temp := server.Blocks
			if temp == nil {
				temp = map[blocks.Block]int{}
			}
			temp[tar.Block] = 1
			server.Blocks = temp
			serverList[i2] = server

			temp = info.Backup
			if temp == nil {
				temp = map[blocks.Block]int{}
			}
			temp[tar.Block] += 1
			info.Backup = temp
			orgInfo[orgName] = info

			i2++
			backupNum++
			// 消息维护
			message, flag := res[server.PeerId]
			if !flag {
				message = bsmsg.New(false)
			}
			message.AddBlock(tar.Block)
			res[server.PeerId] = message
		}
		if backupNum < tar.BackupTar {
			return nil, fmt.Errorf("block%s备份数为%v", tar.Block.Cid().String(), backupNum)
		}
	}
	return res, nil
}

/* 描述: 查询当前文件块在服务节点上的分布情况
 * 入参:[某文件dag顺序的文件块列表，服务节点列表]
 * 出参：[文件各节点备份数量，各文件块需备份数量，各组织文件备份数量]
 */
func findAllocateCondition(blockList []blocks.Block, serverList []standard.CorePeer) ([]BackupResult, []BlockBackupTar, map[string]OrgBackupInfo, error) {
	results := make([]BackupResult, len(serverList))
	peerMap := map[string]int{}
	orgInfo := map[string]OrgBackupInfo{}
	for i, s := range serverList {
		results[i] = BackupResult{
			OrgName: s.Org,
			PeerId:  s.PeerId,
			Blocks:  map[blocks.Block]int{},
		}
		peerMap[s.PeerId] = i
		// 维护组织节点信息
		info, flag := orgInfo[s.Org]
		if flag {
			info = OrgBackupInfo{
				Backup:  nil,
				OrgName: s.Org,
				PeerNum: info.PeerNum + 1,
			}
		} else {
			info = OrgBackupInfo{
				Backup:  nil,
				OrgName: s.Org,
				PeerNum: 1,
			}
		}
		orgInfo[s.Org] = info
	}
	bList := make([]BlockBackupTar, len(blockList))

	// 检索已经存储的情况
	for i, block := range blockList {
		peers := findBlocAllocate(block.Cid())
		validBackupNum := 0
		// 维护已存储信息
		for _, s := range peers {
			i3, flag := peerMap[s]
			if flag {
				validBackupNum++
				b := results[i3].Blocks
				b[block] += 1
				results[i3].Blocks = b
				// 维护组织存储信息
				org := orgInfo[results[i3].OrgName]
				b1 := org.Backup
				if b1 == nil {
					b1 = map[blocks.Block]int{}
				}
				b1[block] += 1
				org.Backup = b1
				orgInfo[results[i3].OrgName] = org
			}
		}
		bList[i] = BlockBackupTar{
			Block:     block,
			BackupTar: DefaultBackupTar - validBackupNum,
		}
	}
	return results, bList, orgInfo, nil
}

/* 描述: 查询某个文件块在ipfs网络的哪些节点有分部
 * 入参:[文件块cid]
 * 出参：[拥有该文件块的节点列表]
 */
func findBlocAllocate(cid cid.Cid) []string {
	if cid.String() == "QmTTA2daxGqo5denp6SwLzzkLJm3fuisYEi9CoWsuHpzfb" {
		return []string{"org0peer0"}
	}
	return nil
}

/* 描述: 分发文件块
 * 入参:[ctx,文件的dag顺序文件块列表，服务节点列表]
 * 后置状态: 经计算后文件块发送到指定服务节点
 */
func (bs *Bitswap) AllocateBlocks(ctx context.Context, results []BackupResult, bList []BlockBackupTar, orgInfo map[string]OrgBackupInfo) error {

	// todo 检查是否满足备份要求

	// 分发算法
	res, err := allocateBlocksAlgorithm(results, bList, orgInfo)
	if err != nil {
		return err
	}

	// 分发文件块
	// todo 多线程分发
	for idString, message := range res {
		go bs.SendBlocks(ctx, idString, message)
	}
	return nil
}
