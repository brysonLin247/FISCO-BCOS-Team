rm *.abi *.bin hello_world.go share_delivery_info.go
# 编译合约
/Users/ismart/Github/go-sdk/solc-0.4.25 --bin --abi -o . ./helloWorld.sol
# 生成.go文件，该文件的有部署合约和与HelloWorld智能合约交互的代码
/Users/ismart/Github/go-sdk/abigen --bin ./helloWorld.bin --abi ./HelloWorld.abi --pkg mycontract --type HelloWorld --out ./hello_world.go

# 编译合约
/Users/ismart/Github/go-sdk/solc-0.4.25 --bin --abi -o . ./ShareDeliveryInfo.sol
# 生成.go文件，该文件的有部署合约和与HelloWorld智能合约交互的代码
/Users/ismart/Github/go-sdk/abigen --bin ./ShareDeliveryInfo.bin --abi ./ShareDeliveryInfo.abi --pkg mycontract --type ShareDeliveryInfo --out ./share_delivery_info.go


rm *.abi *.bin