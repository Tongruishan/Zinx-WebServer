package ziface

//创建接口
type ISever interface {
	Start()
	Stop()
	Sever()
	AddRouter(msgId uint32,router IRouter)
	GetConnMsg()IConnManager
	//注册hook函数
	AddOnConnStart(hookF func(conn IConnection))
	AddOnConnStop(hookF func(conn IConnection))
	//调用hook函数
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}



