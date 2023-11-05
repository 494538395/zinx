package iface

// IServer 服务器通用接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 给 server 添加 Router
	AddRouter(msgID uint32, router Router)
	// GetConnManager 获取连接管理器
	GetConnManager() IConnManager
	// SetOnConnStart 设置 conn start hook
	SetOnConnStart(f func(connection IConnection))
	// SetOnConnClose 设置 conn close hook
	SetOnConnClose(f func(connection IConnection))
	// CallOnConnStart 调用 hook
	CallOnConnStart(conn IConnection)
	// CallOnConnClose 调用 hook
	CallOnConnClose(conn IConnection)
}
