package net

import (
	"fmt"
	"github.com/494538395/zinx/config"
	"github.com/494538395/zinx/iface"
	logger "github.com/494538395/zinx/log"
	"net"
)

//
// Server
//  @Description: 服务器结构
//
type Server struct {
	ServerName  string                             // 服务器名称
	IPVersion   string                             // IP 版本
	IP          string                             // IP
	Port        int                                // 服务器监听端口
	MsgHandler  iface.IMsgHandler                  // 消息处理器
	ConnManager iface.IConnManager                 // 连接管理器
	OnConnStart func(connection iface.IConnection) // hook conn start
	OnConnClose func(connection iface.IConnection) // hook conn closed
}

//
// Start
//  @Description: 服务器启动方法
//
func (s *Server) Start() {
	logger.SetUp()

	logger.Debug(fmt.Sprintf("[Start] Server Listener at IP:%s, Port:%d, is starting", s.IP, s.Port))
	go func() {
		// 开启 workPool
		s.MsgHandler.StartWorkPool()

		// 1.获取 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			logger.Fatal("[Start] resolver tcp addr error", err)
			return
		}
		// 2. 监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			logger.Fatal("[Start] listen tcp error", err)
			return
		}
		logger.Debug(fmt.Sprintf("[Start] start zinx server success,%s,now is listening consistently", s.ServerName))
		var cid int32

		// 3. 阻塞等待客户端连接，处理业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				logger.Debug("[Start] accept error:", err)
			}
			logger.Debug("[Start] 有一个客户端和服务端建立了 tcp 连接")

			if s.ConnManager.GetConnCnt() > int(config.Config.Server.MaxConnSize) {
				// TODO:告知客户端连接关闭
				// 拒绝连接
				logger.Warn("[Server][Start] conn is refused,because conn size exceed limit")
				conn.Close()
				continue
			}

			// 将新的客户端连接和 connection 模块绑定
			connection := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			go connection.Start()
		}
	}()
}

//
// Stop
//  @Description: 服务器停止方法
//
func (s *Server) Stop() {
	// 关闭所有连接
	s.ConnManager.RemoveAllConn()
}

//
// Serve
//  @Description: 服务器运行方法
//
func (s *Server) Serve() {
	// 启动服务器
	s.Start()

	// TODO：这里可以做一些启动服务器之后额外到业务。目的是将启动服务器和其他业务隔离

	// 阻塞进程
	select {}
}

func (s *Server) AddRouter(msgID uint32, router iface.Router) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("[AddRouter] router add")
}

//
// NewServer
//  @Description: 初始化服务器
//  @param name
//  @return iface.IServer
//
func NewServer() iface.IServer {
	return &Server{
		ServerName:  config.Config.Server.ServerName,
		IP:          config.Config.Server.IP,
		Port:        config.Config.Server.Port,
		IPVersion:   config.Config.Server.IPVersion,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
}

func (s *Server) GetConnManager() iface.IConnManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(f func(connection iface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnClose(f func(connection iface.IConnection)) {
	s.OnConnClose = f
}

func (s *Server) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(conn)
		logger.Info("OnConnStart func call,connID:", conn.GetConnID())
	}
}

func (s *Server) CallOnConnClose(conn iface.IConnection) {
	if s.OnConnClose != nil {
		s.OnConnClose(conn)
		logger.Info("OnConnClose func call,connID:", conn.GetConnID())
	}
}
