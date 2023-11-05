package iface

import "net"

type IConnection interface {
	// Start 启动连接
	Start()

	// Stop 停止连接
	Stop()

	// GetTCPConnection 获取当前连接绑定的 socket conn
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前连接模块的连接 ID
	GetConnID() int32

	// RemoteAddr 获取远程客户端的 TCP 状态，IP 和 Port
	RemoteAddr() net.Addr

	// SendMsg 将数据进行封包，然后发送给远程客户端
	SendMsg(msgID uint32, data []byte) error
}

// HandleFunc 处理连接业务的函数
// *net.TCPConn：与客户端的连接
// []byte 处理的数据
// int 处理数据的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
