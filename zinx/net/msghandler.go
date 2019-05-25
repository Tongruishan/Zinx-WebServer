package net

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
)

//存放路由的集合
type MsgHandler struct {
	APIS map[uint32]ziface.IRouter
}

func NewMsgHandler()*MsgHandler{

	return &MsgHandler{
		APIS:make(map[uint32]ziface.IRouter),
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



