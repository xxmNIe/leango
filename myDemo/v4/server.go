package main

import (
	"fmt"
	"leango/zinx/ziface"
	"leango/zinx/znet"
)

type pingRouter struct {
	znet.BaseRouter
}

func (p *pingRouter)PreHandle(request ziface.IRequest){
	fmt.Println(" pre pingRouter.......... ")
	_,err := request.GetConnection().GetConnection().Write([]byte("pre ping...."))
	if err !=nil{
		fmt.Println("call  back pre pingRouter err : ",err)
	}
}
func (p *pingRouter)Handle(request ziface.IRequest){
	fmt.Println("  pingRouter.......... ")
	_,err := request.GetConnection().GetConnection().Write([]byte(" ping...."))
	if err !=nil{
		fmt.Println("call  back  pingRouter err : ",err)
	}
}
func (p *pingRouter)PostHandle(request ziface.IRequest){
	fmt.Println(" post pingRouter.......... ")
	_,err := request.GetConnection().GetConnection().Write([]byte("post ping...."))
	if err !=nil{
		fmt.Println("call  back post pingRouter err : ",err)
	}
}

func main() {
		s:=znet.NewServer("v2")
		s.AddRouter(&pingRouter{})
		s.Server()
}