package ziface

type IMsgHandler interface {
	//添加路由器到map中
	AddMsgRouter(id uint32,router IRouter)
	//调度路由器
	DoMsgHandler(request IRequest)
	//启动任务池
	StartWorkPool()
	//将消息添加到任务池
	SendMsgToTaskSue(request IRequest)

}