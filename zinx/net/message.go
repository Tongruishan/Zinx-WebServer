package net

type Message struct {
	MsgId uint32
	MsgData []byte
	MsgLen uint32

}

func NewMesg(id uint32,data []byte)*Message{

	return &Message{
		MsgId:id,
		MsgData:data,
		MsgLen:uint32(len(data)),
	}

}

func (m *Message)GetMsgLen()uint32{
	return m.MsgLen
}

func (m *Message)GetMsgId()uint32{
	return m.MsgId
}

func (m *Message)GetMsgData()[]byte{
	return m.MsgData
}

func (m *Message)SetMsgLen(len uint32){
	m.MsgLen=len

}
func (m *Message)SetMsgId(id uint32){
	m.MsgId=id

}
func (m *Message)SetMsgData(data []byte){
	m.MsgData=data

}


