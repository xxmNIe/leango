package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"leango/zinx/utils"
	"leango/zinx/ziface"
)

type   DataPack struct {}
/*
格式  Len Id data
     4    4
    首部固定8字节
 */



//获取包长度的方法
func (d *DataPack)GetHeadLen() uint32{

	//Len 4字节     Id 4字节  =8
	return 8
}
//封包方法
func (d *DataPack)Pack(msg ziface.IMessage) ([]byte,error){

	//创建一个存放bytes的buf
	dataBuff := bytes.NewBuffer([]byte{})

	//将data的len id data 依次写入buff中
	if err :=binary.Write(dataBuff,binary.LittleEndian,msg.GetDataLen());err !=nil{
		return nil,err
	}
	if err :=binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId());err !=nil{
		return nil,err
	}
	if err :=binary.Write(dataBuff,binary.LittleEndian,msg.GetData());err !=nil{
		return nil,err
	}

	return dataBuff.Bytes(),nil

}
//拆包方法
func (d *DataPack)UnPack(binaryData []byte) (ziface.IMessage,error){

	dataBuff := bytes.NewBuffer(binaryData)
	msg  :=&Message{}

	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err !=nil{
		return nil,err
	}
	if err:=binary.Read(dataBuff,binary.LittleEndian,&msg.MsgId);err !=nil{
		return nil,err
	}



	if utils.GlobalObject.MaxPackageSize >0 && msg.DataLen > utils.GlobalObject.MaxPackageSize{
		return nil ,errors.New("too Large msg data recv!")
	}

	return msg ,nil
}

func NewDataPack() *DataPack{
	return &DataPack{}
}