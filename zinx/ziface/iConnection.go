package ziface

import "net"

type IConnection interface {
	Start()
	Stop()

	GetConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	SendMsg(uint32,[]byte ) error


	GetProperty(key string)(interface{},error)
	SetProperty(key string,vakue interface{})
	RemoveProperty(key string)

}
//
type  HandleFunc func(*net.TCPConn,[]byte,int ) error