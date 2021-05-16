package znet

import (
	"fmt"
	"leango/zinx/utils"
	"leango/zinx/ziface"
	"strconv"
)

type MsgHandler struct {
	//每个MsgID,对应的处理方法
	Apis map[uint32]ziface.IRouter

	//消息队列，负责worker取任务
	TaskQueue []chan ziface.IRequest
	//worker 工作池大小
	WorkPoolSize uint32
}

func NewMsgHandle() *MsgHandler{
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue: make([]chan ziface.IRequest,utils.GlobalObject.WorkPoolSize),
	}
}


//根据msgId 调度执行对应的Router
func (mh *MsgHandler)DoMsgHandler(requset ziface.IRequest){
	handler ,ok :=mh.Apis[requset.GetMsgId()]
	if !ok {
		fmt.Println("api msgId=",requset.GetMsgId()," not found , need registry")
	}

	handler.PreHandle(requset)
	handler.Handle(requset)
	handler.PostHandle(requset)
}
//为消息添加具体的处理逻辑   router
func (mh *MsgHandler)AddRouter(msgID uint32, router ziface.IRouter){
	//判断重复注册
	if _,ok := mh.Apis[msgID]; ok{
		panic("repeat api, msgID = "+strconv.Itoa(int(msgID)))
	}

	//添加
	mh.Apis[msgID] = router
    fmt.Println("Add Apis msgId=",msgID," succ")
}

//只调用一次
func(mh *MsgHandler) StartWorkerPool(){
	//根据WorkerPoolSize开启
	for i:=0;i<int(mh.WorkPoolSize);i++{
		mh.TaskQueue[i] = make(chan ziface.IRequest,utils.GlobalObject.MaxWorkerTaskLen)

		go mh.StartOneWorker(i,mh.TaskQueue[i])
	}
}

func(mh *MsgHandler) StartOneWorker(workerId int,workerQueue chan ziface.IRequest) {
	fmt.Println("Worker ID=",workerId," started....")
	for{
		select {
		//如果有消息过来,出列就是一个客户端Request,执行当前request锁绑定的业务
			case request := <-workerQueue:
				mh.DoMsgHandler(request)
		}
	}
}


func (mh *MsgHandler) SendMsgToTask(request ziface.IRequest){
		//取余查找worker
		workerId :=request.GetConnection().GetConnID()%mh.WorkPoolSize
		fmt.Println("Add ConnID=",request.GetConnection().GetConnID(),
			" request MsgID=",request.GetMsgId()," to WorkerID=",workerId)
		//送入

		mh.TaskQueue[workerId]<-request
}