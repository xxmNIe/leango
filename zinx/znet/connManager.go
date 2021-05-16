package znet

import (
	"errors"
	"fmt"
	"leango/zinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection

	connLock sync.RWMutex
}

func NewConnManager()*ConnManager{
	return &ConnManager{
		connections: make( map[uint32]ziface.IConnection),
	}
}


//添加连接
func (cm *ConnManager)Add(conn ziface.IConnection){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//将conn加入

	cm.connections[conn.GetConnID()]=conn
	fmt.Println("connID=",conn.GetConnID()," add to ConntionManager succ , connsnums=",cm.Len())
}
//删除连接
func (cm *ConnManager)Remove(conn ziface.IConnection){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections,conn.GetConnID())
	fmt.Println("connID=",conn.GetConnID()," remove from ConntionManager  , connsnums=",cm.Len())
}
//根据connID获取连接
func (cm *ConnManager)Get(connID uint32)(ziface.IConnection,error){
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn,ok:=cm.connections[connID];ok{
		return conn,nil
	}else{
		return nil,errors.New("conn not found !!!")
	}
}


//获取连接总数
func (cm *ConnManager)Len() int{
	return len(cm.connections)
}
//清除并终止连接
func (cm *ConnManager)CleanConn(){
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for connID,conn := range cm.connections{
		//停止
		conn.Stop()
		//删除
		delete(cm.connections,connID)
	}
}