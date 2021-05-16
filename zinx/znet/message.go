package znet

type Message struct {
	DataLen uint32
	MsgId uint32

	Data []byte
}
func NewMessage(msgId uint32,data []byte) *Message{
	return &Message{
		MsgId: msgId,
		DataLen: uint32(len(data)),
		Data: data,
	}
}

func (m *Message)GetMsgId() uint32{
	return m.MsgId
}

func (m *Message)GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message)GetData() []byte {
	return m.Data
}

func (m *Message)SetMsgId(id uint32){
	m.MsgId = id
}

func (m *Message)SetData(data []byte) {
	m.Data = data
}

func (m *Message)SetDataLen(len uint32) {
	m.DataLen = len
}