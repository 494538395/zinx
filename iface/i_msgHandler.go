package iface

type IMsgHandler interface {
	// Handle 处理消息
	Handle(request IRequest)
	// AddRouter 为 msgID 添加 router
	AddRouter(msgID uint32, router Router)

	StartWorkPool()
	// SendMsgToTaskQueue 将消息发送给 队列，异步处理
	SendMsgToTaskQueue(req IRequest)
}
