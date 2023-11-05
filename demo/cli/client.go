package main

import (
	"fmt"
	net2 "github.com/494538395/zinx/net"
	logger "github.com/sirupsen/logrus"
	"io"
	"net"
)

func main() {
	logger.SetLevel(logger.DebugLevel)
	logger.Info("demo show")

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		logger.Error("dial error:", err)
		return
	}

	// 给服务器发送数据
	dp := net2.NewDataPack()

	message := net2.NewMessagePackage(0, []byte("no human is limited"))

	bytes, err := dp.Pack(message)
	if err != nil {
		panic(err)
	}

	n, err := conn.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("conn write n-->", n)

	// 接收服务器的返回
	for {
		// 读取 head 信息
		headData := make([]byte, dp.GetHeadLen())

		cnt, err := io.ReadFull(conn, headData)
		if err != nil {
			panic(err)
		}
		fmt.Println("read head cnt-->", cnt)

		msgHead, err := dp.UnPack(headData)
		if err != nil {
			panic(err)
		}
		// 根据 head 信息读取 data
		if msgHead.GetMsgLen() <= 0 {
			continue
		}
		msg := msgHead.(*net2.Message)
		msg.Data = make([]byte, msgHead.GetMsgLen())
		cnt, err = io.ReadFull(conn, msg.Data)
		if err != nil {
			panic(err)
		}
		fmt.Println("read data cnt-->", cnt)
	}

}
