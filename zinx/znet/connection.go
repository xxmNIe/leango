package znet

import (
	"errors"
	"fmt"
	"io"
	"leango/zinx/utils"
	"leango/zinx/ziface"
	"net"
	"sync"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool


   //告知当前链接已经退出/停止的channel(Reader告知writer)
	ExitChan chan bool

   //无缓冲管道,御用reader和writer通信
	msgChan chan []byte

	MshHandler ziface.IMsgHandler

	TCPServer ziface.IServer

	property map[string]interface{}

	propertylock sync.RWMutex
}

func (c *Connection)Start(){
	fmt.Println("Conn start() ... ConnID = ",c.ConnID)
	//启动从当链接读取数据的业务

	go c.StartReader()

	go c.StartWriter()

	//按照开发者传递进来的创建后需要调用的业务,执行对应的hook
	c.TCPServer.CallOnConnStart(c)

}
func (c *Connection)Stop(){

	fmt.Println("Conn Sttop(),,, ConnID=",c.ConnID)
	if c.isClosed==true{
		return
	}


	//调用开发者传进来的销毁之前的hook函数
	c.TCPServer.CallOnConnStop(c)

	//关闭原始连接
	c.Conn.Close()

	c.isClosed = true
	close(c.ExitChan)

	close(c.msgChan)

	//当前连接错connmager中摘除
	c.TCPServer.GetConnMgr().Remove(c)
	
}

func (c *Connection)GetConnection() *net.TCPConn{
	return c.Conn
}

func (c *Connection)GetConnID() uint32{
	return  c.ConnID
}

func (c *Connection)RemoteAddr() net.Addr{
	addr := c.Conn.RemoteAddr()
	return addr
}

func (c* Connection)SendMsg(msgId uint32,data []byte ) error{

	if c.isClosed==true {
		return errors.New("conn is closed")
	}
	dp :=&DataPack{}
	binarydata,err := dp.Pack(NewMessage(msgId,data))
	if err!=nil{
		fmt.Println("pack error: ",err)
	}
	//发送数据
	 c.msgChan <- binarydata


	return nil
}

//链接的读业务
func (c *Connection) StartReader() {
	fmt.Println("[Reader] Goroutine is running")
	cl:= c.RemoteAddr().String()
	defer fmt.Println("Conn=",c.ConnID," [Reader is exit],remote addr is ",cl)
	defer c.Stop()

	for{

		dp :=&DataPack{}
		headData := make([]byte,dp.GetHeadLen())

		if _,err :=io.ReadFull(c.GetConnection(),headData);err!=nil{
			fmt.Println("Read msg error: ",err)
			break
		}
		msg ,err:= dp.UnPack(headData)
		if err !=nil{
			fmt.Println("unpack error: ",err)
			break
		}
		var data []byte
		if msg.GetDataLen() >0{
			data = make([]byte,msg.GetDataLen())
			if _,err :=io.ReadFull(c.GetConnection(),data);err !=nil{
				fmt.Println("read data error: ",err)
				break

			}
		}
		msg.SetData(data)

		req :=Request{
			conn: c,
			msg: msg,
		}
		if utils.GlobalObject.WorkPoolSize >0{
			c.MshHandler.SendMsgToTask(&req)
		}else {
			go c.MshHandler.DoMsgHandler(&req)
		}

	}
}

func (c *Connection)StartWriter(){
	fmt.Println("[Writer Gortine is running]")

	defer fmt.Println(c.RemoteAddr().String()," [conn Writer exit]")

	for{
		select {
			case data :=<-c.msgChan:
				if _,err := c.Conn.Write(data);err!=nil{
					fmt.Println("Send data err:",err)
					return
				}
			case <-c.ExitChan:
				//raeader 退出
				return
		}
	}
}


func (c *Connection) GetProperty(key string) (interface{},error) {
	c.propertylock.RLock()
	defer c.propertylock.RUnlock()
	if value,ok :=c.property[key];!ok{
		return value,nil
	}
	return nil,errors.New("property not found")

}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertylock.Lock()
	defer c.propertylock.Unlock()
	c.property[key] = value

}

func (c *Connection) RemoveProperty(key string) {
	c.propertylock.Lock()
	defer c.propertylock.Unlock()
	delete(c.property,key)
}

func NewConnection(server ziface.IServer,conn *net.TCPConn,connID uint32,msghandler ziface.IMsgHandler) *Connection{
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		MshHandler: msghandler,
		msgChan: make(chan []byte),
		isClosed: false,
		ExitChan: make(chan bool,1),
		TCPServer:server,
		property: make(map[string]interface{}),
	}
	c.TCPServer.GetConnMgr().Add(c)
	return c
}