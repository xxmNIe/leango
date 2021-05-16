package main

import (
	"fmt"
	"leango/zinx/ziface"
	"leango/zinx/znet"
)

type pingRouter struct {
	znet.BaseRouter
}

func DoConnBegin(conn ziface.IConnection){
	fmt.Println("DoConnBegin Called.....")
	if err:= conn.SendMsg(202,[]byte("DoConnBegin"));err!=nil{
		fmt.Println(err)
	}
}

func DoConnEnd(conn ziface.IConnection){
	fmt.Println("DoConnEnd Called.....")
	if err:= conn.SendMsg(202,[]byte("DoConnEnd"));err!=nil{
		fmt.Println(err)
	}
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
		s.SetOnConnStart(DoConnBegin)
		s.SetOnConnStop(DoConnEnd)
		s.Server()
}