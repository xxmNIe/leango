package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	//给当前服务提供一个router,供处理客户端链接使用
	AddRouter(msgID uint32,router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)



}