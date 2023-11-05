package net

import (
	"errors"
	"fmt"
	"github.com/494538395/zinx/iface"
	logger "github.com/494538395/zinx/log"
	"io"
	"net"
)

// Connection 连接模块
type Connection struct {
	// Server connection 所属服务器
	Server iface.IServer

	// 当前连接的 socket TCP 套接字
	Conn *net.TCPConn

	// 连接 ID
	ConnID int32

	// 无缓冲 channel，用于 Read、Write 之间的通信
	send chan []byte

	// 当前连接的状态
	isClosed bool

	// 告知当前连接已经退出的 channel
	ExitChan chan bool

	// 改连接处理的的方法
	MasHandler iface.IMsgHandler
}

// NewConnection 初始化连接模块
func NewConnection(server iface.IServer, conn *net.TCPConn, connID int32, msgHandler iface.IMsgHandler) *Connection {
	c := &Connection{
		Server:     server,
		Conn:       conn,
		ConnID:     connID,
		ExitChan:   make(chan bool, 1),
		send:       make(chan []byte),
		MasHandler: msgHandler,
	}

	c.Server.GetConnManager().AddConn(c)

	return c
}

func (c *Connection) Start() {
	logger.Debug("Conn Start ConnID:", c.ConnID)
	// 开启两个协程，两个协程分别负责从客户端读取数据、向客户端写数据
	go c.Read()
	go c.Write()
	// call conn start hook
	c.Server.CallOnConnStart(c)
}

// Read 接收客户端消息
func (c *Connection) Read() {
	logger.Debug("[Connection] [Read] ConnID:", c.ConnID)
	defer func() {
		logger.Debug("[Connection] [Read] ConnID:", c.ConnID, "Read is exit, remote addr is ", c.RemoteAddr().String())
		c.Stop()
	}()

	for {
		// 字节缓冲区，这里的意思是，最多一次读取 512 字节长度的数据到缓冲区,即 cnt <= 512
		//buf := make([]byte, 512)
		//cnt, err := c.Conn.Read(buf)
		//fmt.Println(123)
		//if err != nil {
		//	if err == io.EOF {
		//		// 客户端连接已经断开
		//		logger.Warn("ConnID:", c.ConnID, " conn is break:", err)
		//		break
		//	}
		//	logger.Error("recv buf error:", err)
		//	continue
		//}
		// 得到当前连接的 request
		// 拆包
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		cnt, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			logger.Error("read full error:", err)
			return
		}
		logger.Info("read head data cnt:", cnt)

		// 得到包头数据
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			logger.Error("unpack error:", err)
			return
		}
		if msgHead.GetMsgLen() <= 0 {
			continue
		}

		// 读取 data 数据
		msg := msgHead.(*Message)
		msg.Data = make([]byte, msgHead.GetMsgLen())

		cnt, err = io.ReadFull(c.Conn, msg.Data)
		if err != nil {
			logger.Error("read full error:", err)
			return
		}
		logger.Info("read data cnt:", cnt)

		req := &Request{conn: c, msg: msg}
		// 将 req 发送给 taskQueue ，异步处理，而不是每来一条客户端消息都开一个 goroutine
		//go c.MasHandler.Handle(req)
		c.MasHandler.SendMsgToTaskQueue(req)

		logger.Debug("read continue")
	}
}

// Write 给客户端发送消息
func (c *Connection) Write() {
	logger.Debug("[Connection] [Write] ConnID:", c.ConnID)
	defer func() {
		logger.Debug("[Connection] [Write] ConnID:", c.ConnID, "Write is exit, remote addr is ", c.RemoteAddr().String())
		c.Stop()
	}()

	for {

		select {
		case msg, ok := <-c.send:
			if !ok {
				return
			}
			// 给客户端发送消息
			cnt, err := c.Conn.Write(msg)
			if err != nil {
				logger.Fatal("[Connection] [Write] conn write error:", err)
			}
			logger.Debug("[Connection] [Write] conn write cnt:", cnt)
		case <-c.ExitChan:
			// 走到这里，说明 Read 已经退出，则 Write 也退出
			return
		}
	}

}

func (c *Connection) Stop() {
	logger.Debug("Conn Stop ConnID:", c.ConnID)
	if c.isClosed {
		return
	}
	// call conn closed hook
	c.Server.CallOnConnClose(c)

	c.Conn.Close()
	c.isClosed = true
	close(c.ExitChan)
	close(c.send)
	c.Server.GetConnManager().RemoveConn(c)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() int32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 将要发送给客户端的数据先进行封包，然后发送
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("conn is closed,can not send msg")
	}
	// 封包
	dp := NewDataPack()
	bytes, err := dp.Pack(NewMessagePackage(msgID, data))
	if err != nil {
		logger.Error("pack message error:", err)
		return errors.New(fmt.Sprintf("pack message error:%v", err))
	}

	// 发送给客户端
	c.send <- bytes

	return nil
}
