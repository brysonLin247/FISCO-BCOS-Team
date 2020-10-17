package main

import (
	"ShareDeliveryInfo/smartcontract"
	"ShareDeliveryInfo/socket"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func main() {
	config := &conf.ParseConfig("conf/config.toml")[0]

	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	//// 部署合约
	//contractAddress,transaction, ShareLogisticInfoInstance,err := mycontract.DeployHelloWorld(client.GetTransactOpts(),client)
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("HelloWorld 合约地址：",contractAddress.Hex())
	//fmt.Println("部署HelloWorld 合约的交易hash",transaction.Hash().Hex())

	// 调用已部署的合约
	// contractAddress := common.HexToAddress("0xc1bb692e753f47577a96A08ac642e6E38b1A8B87")
	// ShareLogisticInfoInstance, err := mycontract.NewHelloWorld(contractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 跟合约交互的Session
	// ShareLogisticInfoSession := mycontract.HelloWorldSession{
	// 	Contract: ShareLogisticInfoInstance,
	// 	CallOpts:*client.GetCallOpts(),
	// 	TransactOpts:*client.GetTransactOpts(),
	// }

	// // 创建TCP Server，接收RFID上传的信息，并上传到链上
	// socket.ServerSocket(ShareLogisticInfoSession, client)


	// 调用已部署的ShareDeliveryInfo合约
	contractAddress := common.HexToAddress("0xdd98a9b4cff830d48c8242111eb63660cd84ecca")
	ShareLogisticInfoInstance, err := mycontract.NewShareDeliveryInfo(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 跟合约交互的Session
	ShareLogisticInfoSession := mycontract.ShareDeliveryInfoSession{
		Contract: ShareLogisticInfoInstance,
		CallOpts:*client.GetCallOpts(),
		TransactOpts:*client.GetTransactOpts(),
	}

	// 创建TCP Server，接收RFID上传的信息，并上传到链上
	socket.ServerSocket(ShareLogisticInfoSession, client)

}
