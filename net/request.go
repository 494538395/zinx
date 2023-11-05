package net

import "github.com/494538395/zinx/iface"

type Request struct {
	// 已经和客户端完成建立连接到连接
	conn iface.IConnection
	// 客户端到请求数据
	msg iface.IMessage
}

func (r *Request) GetConn() iface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

func (r *Request) GetMsgLen() uint32 {
	return r.msg.GetMsgLen()
}
