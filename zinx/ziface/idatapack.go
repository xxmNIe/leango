package ziface

//按照协议去解析或者封装包
type IDataPack interface {
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte,error)
	//拆包方法
	UnPack(data []byte) (IMessage,error)
}