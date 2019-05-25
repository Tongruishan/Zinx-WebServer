package ziface

//创建接口
type ISever interface {
	Start()
	Stop()
	Sever()
	AddRouter(msgId uint32,router IRouter)
}



