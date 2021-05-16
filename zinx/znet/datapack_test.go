package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)
//拆包封包单元测试
func TestDataPack(t *testing.T){


	//
	listenner,err :=net.Listen("tcp","127.0.0.1:7777")
	if err!=nil{
		fmt.Println("Server listenner err:",err)
		return
	}
	//
	go func() {
		//
		for{
			conn,err := listenner.Accept()
			if err!=nil{
				fmt.Println("server accept err :",err)
			}
			go func(conn net.Conn) {
				//unpack
				dp :=&DataPack{}
				for{
					//读两次
					//先读头
					headData :=make([]byte,dp.GetHeadLen())
					_,err:= io.ReadFull(conn,headData)
					if err!=nil{
						fmt.Println("read  head err :",err)
						return
					}

					msgHead,err :=dp.UnPack(headData)
					if err!=nil{
						fmt.Println("server unpack err :",err)
						return
					}
					//如果有数据，第二次读
					if msgHead.GetDataLen() >0{
						//msg has data

						msg :=msgHead.(*Message)
						msg.Data = make([]byte,msg.GetDataLen())

						_,err :=io.ReadFull(conn,msg.Data)
						if err!=nil{
							fmt.Println("server unpack  read data err :",err)
							return
						}

						fmt.Println("--->Recv  MsgId: ",msg.MsgId," MsgData len :",msg.DataLen," Msgdata: " ,string(msg.Data))
					}


				}
			}(conn)
		}
	}()

	conn ,_:=net.Dial("tcp","127.0.0.1:7777")
	dp := &DataPack{}



	//封装两个包发送

	msg1 :=&Message{
		MsgId: 1,
		DataLen: 4,
		Data: []byte{'Z','i','n','x'},
	}

	data1,err := dp.Pack(msg1)
	if err!=nil{
		fmt.Println("client pack err",err)
	}

	msg2 :=&Message{
		MsgId: 2,
		DataLen: 5,
		Data: []byte{'n','i','h','a','o'},
	}

	data2,err := dp.Pack(msg2)
	if err!=nil{
		fmt.Println("client pack err",err)
	}

	data1 = append(data1,data2...)

	go func() {

		for {
			conn.Write(data1)
			time.Sleep(1 * time.Second)
		}
	}()


	select {

	}
}
