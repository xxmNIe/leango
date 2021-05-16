package ziface

//实际上吧客户端请求的链接信息和请求数据封装

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
	GetDatalen() uint32
	GetMsgId() uint32
}