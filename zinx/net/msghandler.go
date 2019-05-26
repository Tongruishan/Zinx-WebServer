package net

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
	"ZinxHouse/Zinx-WebServer/zinx/config"
)

//存放路由的集合
type MsgHandler struct {
	APIS map[uint32]ziface.IRouter


	//消息队列
	TaskQueue []chan ziface.IRequest

	//增加任务池数量
	WorkPoolSize uint32


}

func NewMsgHandler()*MsgHandler{

	return &MsgHandler{
		APIS:make(map[uint32]ziface.IRouter),
		TaskQueue:make([]chan ziface.IRequest,config.GlobleConf.WorkerPoolSize),
		WorkPoolSize:config.GlobleConf.WorkerPoolSize,
	}

}

//添加路由器到map中
func (mh *MsgHandler)AddMsgRouter(msgId uint32,router ziface.IRouter){

	//判断是否添是否有元素
	_,ok:=mh.APIS[msgId]

	if ok{
		fmt.Println("API msgId=",msgId)
		return
	}

	//添加到map中
	mh.APIS[msgId]=router

	fmt.Println("Append Apis msgId=",msgId," successfull")

}
//调度路由器
func (mh *MsgHandler)DoMsgHandler(request ziface.IRequest){

	msgid:=request.GetMsg().GetMsgId()

	router,ok:=mh.APIS[msgid]
	if !ok{
		fmt.Println("router not found!!,msgId=",msgid,"should append!!!")
		return
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)

}

func(mh *MsgHandler)startOnePool(workId int,taskQueue chan ziface.IRequest){

	for{
		select {
		case req:=<-taskQueue:
			mh.DoMsgHandler(req)

		}
	}

}

//启动任务池
func (mh *MsgHandler)StartWorkPool(){
	fmt.Println("Work Pool is Start ......")

	for i:=0;i<int(config.GlobleConf.WorkerPoolSize);i++{

		mh.TaskQueue[i]=make(chan ziface.IRequest,config.GlobleConf.MaxWorkerTaskLen)
		//等待消息被穿过来
		go mh.startOnePool(i,mh.TaskQueue[i])

	}

}
//将消息添加到任务池
func (mh *MsgHandler)SendMsgToTaskSue(request ziface.IRequest){

	workId:=request.GetConn().GetConnId()%config.GlobleConf.WorkerPoolSize

	mh.TaskQueue[workId]<-request
}



