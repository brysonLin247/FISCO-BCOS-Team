package socket

import (
    "strconv"
    "fmt"
    "log"
    "net"
    "ShareDeliveryInfo/mycontract"
    "github.com/FISCO-BCOS/go-sdk/client"
    // "strings"
)

/**
  Server接收上传的RFID信息，然后与合约交互将信息上传到链上。
 */

func connHandler(c net.Conn, session mycontract.ShareDeliveryInfoSession, client *client.Client) {
    //1.conn是否有效
    if c == nil {
        log.Panic("无效的 socket 连接")
    }

    //2.新建网络数据流存储结构
    buf := make([]byte, 4096)
    var bytes []byte

    //3.循环读取网络数据流
    for {
        //3.1 网络数据流读入 buffer
        cnt, err := c.Read(buf)
        //3.2 数据读尽、读取错误 关闭 socket 连接
        if cnt == 0 || err != nil {
            c.Close()
            break
        }

        currBytes := buf[0:cnt]
        bytes = append(bytes, currBytes...)
    }

    // //3.3 根据输入流进行逻辑处理

    originStr := string(bytes)
    fmt.Println("输入字符："+originStr)

    fmt.Printf("包裹%s当前到达的站点: %s\n ",string(bytes[9:12]), string(bytes[24:33]))

    package_uid, err := strconv.ParseUint(string(bytes[9:12]),16,64)
    if err!=nil{
      log.Fatal(err)
    }
    station := string(bytes[24:33])

    // package_uid := uint64(1)
    // station := "test station"

    fmt.Printf("站点读卡器读取包裹编号->%d, 到达站点: %s \n", package_uid, station)
    // c.Write([]byte("服务器端回复" + originStr + "\n"))

    setTransaction,err := session.Set(package_uid, station)
    if err!=nil {
      log.Fatal(err)
    }
    // 等待交易被打包处理
    receipt, err := client.WaitMined(setTransaction)
    if err != nil {
      log.Fatalf("tx mining error:%v\n", err)
    }
    fmt.Printf("信息已上链，获得的结果收据: %s\n", receipt.GetTransactionHash())

    s, err := session.Get(package_uid)
    if err!=nil{
      log.Fatal(err)
    }

    fmt.Printf("包裹%d: %s\n ",package_uid, s)

    //c.Close() //关闭client端的连接，telnet 被强制关闭

    // nextStation, err := session.NextStation(s)

    // fmt.Printf("下一步物流站点是：xxxxx",)

}

//开启serverSocket
func ServerSocket(session mycontract.ShareDeliveryInfoSession, client *client.Client) {
    //1.监听端口
    server, err := net.Listen("tcp", ":1234")

    if err != nil {
        fmt.Println("开启socket服务失败")
    }

    fmt.Println("开启 Server ...")

    for {
        //2.接收来自 client 的连接,会阻塞
        conn, err := server.Accept()

        if err != nil {
            fmt.Println("连接出错")
        }

        //并发模式 接收来自客户端的连接请求，一个连接 建立一个 conn，服务器资源有可能耗尽 BIO模式
        go connHandler(conn, session, client)
    }

}