package main

import (
	"fmt"
	"github.com/494538395/zinx/config"
	"github.com/494538395/zinx/iface"
	"github.com/494538395/zinx/net"
)

type myRouterMsg01 struct {
	net.BaseRouter
}

func (mr *myRouterMsg01) Handle(req iface.IRequest) {
	fmt.Println("myRouterMsg01  处理的消息ID是-->", req.GetMsgID(), "处理的消息内容是--->", string(req.GetData()))
	// 封包，回写
	err := req.GetConn().SendMsg(1, []byte("你好啊"))
	if err != nil {
		panic(err)
	}
}

type myRouterMsg02 struct {
	net.BaseRouter
}

func (mr *myRouterMsg02) Handle(req iface.IRequest) {
	fmt.Println("myRouterMsg02  处理的消息ID是-->", req.GetMsgID(), "处理的消息内容是--->", string(req.GetData()))
	// 封包，回写
	err := req.GetConn().SendMsg(2, []byte("你好啊"))
	if err != nil {
		panic(err)
	}
}

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	server := net.NewServer()
	server.AddRouter(0, &myRouterMsg01{})
	server.AddRouter(1, &myRouterMsg02{})
	server.SetOnConnStart(func(c iface.IConnection) {
		fmt.Println("connID:", c.GetConnID(), "上线了奥！")
	})
	server.SetOnConnClose(func(c iface.IConnection) {
		fmt.Println("connID:", c.GetConnID(), "下线了奥！")
	})

	server.Serve()

}
