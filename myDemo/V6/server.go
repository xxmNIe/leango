package main

import (
	"fmt"
	"leango/zinx/ziface"
	"leango/zinx/znet"
)

type pingRouter struct {
	znet.BaseRouter
}



func (p *pingRouter)Handle(request ziface.IRequest){
	fmt.Println("  Call Router  Handle.......... ")
	fmt.Println("recv from client : msgID=",request.GetMsgId(),
		", data=",string(request.GetData()))

	if err := request.GetConnection().SendMsg(1,[]byte("ping...ping....ping..."));err!=nil{
		fmt.Println("handle send err")

	}
}

type HelloRouter struct {
	znet.BaseRouter
}



func (p *HelloRouter)Handle(request ziface.IRequest){
	fmt.Println("  Call helloRouter  Handle.......... ")
	fmt.Println("recv from client : msgID=",request.GetMsgId(),
		", data=",string(request.GetData()))

	if err := request.GetConnection().SendMsg(1,[]byte("hello...hello....hello..."));err!=nil{
		fmt.Println("handle send err")
	}
}


func main() {
		s:=znet.NewServer("v10000")
		s.AddRouter(0,&pingRouter{})
		s.AddRouter(1,&HelloRouter{})
		s.Server()
}