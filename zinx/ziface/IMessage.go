package ziface

type IMessage interface {
	//获取
	GetMsgLen()uint32
	GetMsgId()uint32
	GetMsgData()[]byte
	//设置
	SetMsgLen(uint32)
	SetMsgId(uint32)
	SetMsgData([]byte)
}