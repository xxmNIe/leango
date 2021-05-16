package main

import (
	"fmt"
	"io"
	"leango/zinx/znet"
	"net"
	"time"
)

func main() {
	fmt.Println("start clinet1....")
	time.Sleep(1*time.Second)
	conn,err :=net.Dial("tcp","127.0.0.1:7777")
	if err !=nil{
		fmt.Println("client start err ,exit")
		return
	}

	for{
		dp :=&znet.DataPack{}

		binarydata,err := dp.Pack(znet.NewMessage(1,[]byte("zinx1V0.6  Test message")))
		if err!=nil{
			fmt.Println("pack error :",err)
		}
		_,err =conn.Write(binarydata)
		if err!=nil{
			fmt.Println("write error: ",err)
			break
		}

		headMsg :=make([]byte,dp.GetHeadLen())

		_,err= io.ReadFull(conn,headMsg)

		if err!=nil{
			fmt.Println("read msg head error: ",err)
		}
		msg ,err := dp.UnPack(headMsg)
		if err!=nil{
			fmt.Println("unpack error: ",err)
			break
		}

		if msg.GetDataLen()>0{
			data :=make([]byte,msg.GetDataLen())
			_,err =io.ReadFull(conn,data)
			if err !=nil{
				fmt.Println("read data error: ",err)
				break
			}
			msg.SetData(data)
		}
		fmt.Println("Recv Server Msg : ID=",msg.GetMsgId(),", data=",string(msg.GetData()))
		time.Sleep(2*time.Second)
	}

}