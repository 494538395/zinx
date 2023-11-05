package net

import (
	"github.com/494538395/zinx/config"
	"github.com/494538395/zinx/iface"
	logger "github.com/494538395/zinx/log"
	"strconv"
)

type MsgHandler struct {
	// msgID 和 router 的映射
	routers map[uint32]iface.Router
	// 工作队列
	TaskQueue []chan iface.IRequest
	// worker 数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		routers:        map[uint32]iface.Router{},
		TaskQueue:      make([]chan iface.IRequest, config.Config.Server.WorkerPoolSize),
		WorkerPoolSize: config.Config.Server.WorkerPoolSize,
	}
}

func (mh *MsgHandler) Handle(request iface.IRequest) {
	router, found := mh.routers[request.GetMsgID()]
	if !found {
		panic("not found msg router,msgID:" + strconv.Itoa(int(request.GetMsgLen())))
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)

}

func (mh *MsgHandler) AddRouter(msgID uint32, router iface.Router) {
	if _, found := mh.routers[msgID]; found {
		panic("repeated msg router,msgID:" + strconv.Itoa(int(msgID)))
	}
	mh.routers[msgID] = router
}

func (mh *MsgHandler) StartWorkPool() {
	logger.Debug("[MsgHandler] StartWorkPool starting")
	defer logger.Debug("[MsgHandler] StartWorkPool finished")

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan iface.IRequest, config.Config.Server.MaxWorkerTaskLen)
		// 启动当前 worker
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandler) StartOneWorker(workID int, taskQueue chan iface.IRequest) {
	logger.Debug("workID:", workID, "started StartOneWorker")

	for {
		select {
		// 出列的就是客户端的 request
		case req, ok := <-taskQueue:
			if !ok {
				return
			}
			mh.Handle(req)
		}
	}
}

// SendMsgToTaskQueue 将消息发送给 队列，异步处理
func (mh *MsgHandler) SendMsgToTaskQueue(req iface.IRequest) {
	// 1.将消息平均分配给不同的 worker
	workerID := req.GetConn().GetConnID() % int32(mh.WorkerPoolSize)

	// 2.消息发送给对应的 队列
	mh.TaskQueue[workerID] <- req
}
