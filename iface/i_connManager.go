package iface

// IConnManager 连接管理器
type IConnManager interface {
	// AddConn 添加新的连接
	AddConn(connection IConnection)
	// RemoveConn 删除连接
	RemoveConn(connection IConnection)
	// GetConn 获取连接
	GetConn(connID uint32) (IConnection, error)
	// GetConnCnt 获取连接个数
	GetConnCnt() int
	// RemoveAllConn 删除所有连接
	RemoveAllConn()
}
