package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelID   = "mychannel"
	chaincodeID = "econtract"
	orgID       = "Org1"
	userID      = "Admin"
)

func main() {
	configOpt := config.FromFile("./config/sdk-config.yaml")
	sdk, err := fabsdk.New(configOpt)
	if err != nil {
		log.Fatalf("创建新的SDK失败: %v\n", err)
		return
	}
	defer sdk.Close()
	log.Printf("---> 创建SDK成功\n")

	var options_user fabsdk.ContextOption
	var options_org fabsdk.ContextOption

	options_user = fabsdk.WithUser(userID)
	options_org = fabsdk.WithOrg(orgID)

	clientChannelContext := sdk.ChannelContext(channelID, options_user, options_org)
	client, err := ledger.New(clientChannelContext)
	if err != nil {
		log.Fatalf("创建sdk客户端失败: %v\n", err)
		return
	}

	info, err := client.QueryInfo()
	if err != nil {
		log.Fatalf("查询区块链概况: %v\n", err)
		return
	}

	fmt.Printf("区块高度:\n%v\n", info.BCI.Height)
	fmt.Printf("当前区块Hash:\n%v\n", hex.EncodeToString(info.BCI.CurrentBlockHash))
	fmt.Printf("前一区块Hash:\n%v\n", hex.EncodeToString(info.BCI.PreviousBlockHash))

	// -------------------- 第1种方式： 根据哈希查询区块 ----------------
	block, err := client.QueryBlockByHash(info.BCI.CurrentBlockHash)
	if err != nil {
		log.Fatalf("查询区块信息失败: %v\n", err)
		return
	}
	fmt.Printf("区块编号: %v\n", block.Header.Number)
	fmt.Printf("区块Hash:\n%v\n", hex.EncodeToString(block.Header.DataHash))

	// -------------------- 第1种方式： 根据块号查询区块 ----------------
	blockNumber := info.BCI.Height - 1
	block, err = client.QueryBlock(blockNumber)

	fmt.Printf("区块编号: %v\n", block.Header.Number)
	fmt.Printf("区块Hash:\n%v\n", hex.EncodeToString(block.Header.DataHash))

	cfg, err := client.QueryConfig()
	fmt.Printf("通道名称: %v\n", cfg.ID())
	fmt.Printf("区块个数: %v\n", cfg.BlockNumber())
	fmt.Printf("锚节点:\n  主机:%v\n  端口:%v\n  机构:%v\n", cfg.AnchorPeers()[0].Host, cfg.AnchorPeers()[0].Port, cfg.AnchorPeers()[0].Org)
}
