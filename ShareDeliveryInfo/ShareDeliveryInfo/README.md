# 编译智能合约

## 版本需求：
abigen版本：1.9.16-stable
solc版本：0.4.25


# 运行服务器端，实现跟智能合约交互
1. 服务器安装Go1.13版本
2. 打开一个终端，进入ShareDeliveryInfo目录，运行go run main.go
运行成功的话，将显示“开启 Server ...”


# 测试数据上链的流程：

开启服务器端后，打开一个新的终端，进入ShareDeliveryInfo目录，运行go run client_main.go，运行成功的话，将显示“请输入客户端请求数据...“,可以输入上传的信息，模拟上传RFID读卡信息。

