package iface

// IRequest 接口。用于将客户端的请求的连接信息和数据信息封装到一起
type IRequest interface {
	// GetConn 得到当前连接
	GetConn() IConnection

	// GetData 得到请求的消息数据
	GetData() []byte

	GetMsgID() uint32

	GetMsgLen() uint32
}
