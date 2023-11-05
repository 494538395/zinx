package iface

// Router 路由抽象接口
type Router interface {
	// PreHandle 处理业务之前到 hook
	PreHandle(req IRequest)
	// Handle 处理业务之的方法
	Handle(req IRequest)
	// PostHandle 处理业务之后的 hook
	PostHandle(req IRequest)
}
