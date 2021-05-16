package znet

import "leango/zinx/ziface"

//实现router时，先嵌入这个基类,根据需求重写方法
//类似抽象类的作用
type BaseRouter struct {
}
//都为空，之后的实现类根据业务需求重写需要的方法即可
func (br *BaseRouter)PreHandle(request ziface.IRequest){}

func (br *BaseRouter)Handle(request ziface.IRequest){}
func (br *BaseRouter)PostHandle(request ziface.IRequest){}

