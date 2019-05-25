package net

import "ZinxHouse/Zinx-WebServer/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	msg ziface.IMessage
}

//初始话对象
func NewRequest(conn ziface.IConnection,msg ziface.IMessage)*Request{

	res:=&Request{
		conn:conn,
		msg:msg,
	}

	return res

}



//获取链接
func(r *Request)GetConn() ziface.IConnection{

	return r.conn

}
//获取数据
func(r *Request)GetMsg()ziface.IMessage{

	return r.msg

}

