package ziface

/*
消息管理抽象层
 */

type IMsgHandler interface {
	//根据msgId 调度执行对应的Router
	DoMsgHandler(requset IRequest)
	//为消息添加具体的处理逻辑   router
	AddRouter(msgID uint32, router IRouter)
	//启动工作池
	StartWorkerPool()

	SendMsgToTask(request IRequest)
}