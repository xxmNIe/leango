package znet

import "leango/zinx/ziface"

type Request struct {
	//客户端的链接
	conn ziface.IConnection
	//客户端请求的数据
	msg ziface.IMessage
}

func(r *Request)GetConnection()ziface.IConnection {
	return r.conn
}
func(r *Request)GetData()[]byte{
	return r.msg.GetData()
}

func(r *Request)GetDatalen() uint32{
	return r.msg.GetDataLen()
}
func(r *Request)GetMsgId() uint32{
	return r.msg.GetMsgId()
}