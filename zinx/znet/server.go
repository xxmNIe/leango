package znet

import (
	"fmt"
	"leango/zinx/utils"
	"leango/zinx/ziface"
	"net"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
	//当前server的消息管理模块,用来绑定msgId和对应的处理业务
	MsgHandler ziface.IMsgHandler

	//连接管理模块

	ConnMgr ziface.IConnManager

	//创建连接之后自动调用的方法
	OnConnStart func(conn ziface.IConnection)
	//销毁连接之前调用的方法
	OnConnStop func(conn ziface.IConnection)




}


func (s *Server)AddRouter(msgID uint32,router ziface.IRouter){
	s.MsgHandler.AddRouter(msgID,router)
	fmt.Println("Add Router Succ")
}



func (s *Server) Start(){
	fmt.Printf("[Zinx] Server name :%s listenner at Ip:%s port:%d is starting \n",utils.GlobalObject.Name,utils.GlobalObject.Host,utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] MaxConn %d,MaxPackageSize:%d\n",utils.GlobalObject.MaxConn,utils.GlobalObject.MaxPackageSize)
	go func() {

		s.MsgHandler.StartWorkerPool()

		taddr,err := net.ResolveTCPAddr("tcp",fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err !=nil{
			fmt.Println("resolve tcp addr err: ",err )
		}

		//监听服务器地址
		listener,err := net.ListenTCP(s.IPVersion,taddr)
		if err !=nil{
			fmt.Println("listen tcp addr err: ",err )
		}
		fmt.Println("start Server ",s.Name," succ")

		var cid uint32
		cid = 0
		//等待客户端链接,处理
		for{
			conn ,err := listener.AcceptTCP()
			if err !=nil{
				fmt.Println("accept err: ",err)
			}
			//判断是否超过最大连接

			if s.ConnMgr.Len() >utils.GlobalObject.MaxConn{
				fmt.Println("Too many   Conn nums , max is ",utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			fmt.Println("              max:",utils.GlobalObject.MaxConn,"                                    conn  nums=",s.ConnMgr.Len())

			fmt.Println(" handler conn ",conn.RemoteAddr())
			//这里把router与conn绑定,所以要在server.server()调用之前添加router
			c:= NewConnection(s,conn,cid,s.MsgHandler)
			cid++
			go c.Start()
		}
	}()

}
func (s *Server) Stop(){

	fmt.Println("[STOP] server name=",s.Name)
	s.ConnMgr.CleanConn()


}
func (s *Server) Server(){
	s.Start()


	select {

	}
}

func ( s *Server)GetConnMgr() ziface.IConnManager{
	return s.ConnMgr
}

func (s *Server)SetOnConnStart(hookfunc func(conn ziface.IConnection)){
 	s.OnConnStart = hookfunc
}
func (s *Server)SetOnConnStop(hookfunc func(conn ziface.IConnection)){
	s.OnConnStop= hookfunc

}
func (s *Server)CallOnConnStart(conn  ziface.IConnection){
 if s.OnConnStart !=nil{
 	fmt.Println("CallOnConnStart Func Call.....")
 	s.OnConnStart(conn)
 }
}
func (s *Server)CallOnConnStop(conn ziface.IConnection){
	if s.OnConnStop !=nil{
		fmt.Println("CallOnConnStop Func Call.....")
		s.OnConnStop(conn)
	}
}



func NewServer(name string) ziface.IServer {
	s:= &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr: NewConnManager(),
	}
	return s
}