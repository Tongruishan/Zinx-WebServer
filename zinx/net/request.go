package net

import "ZinxHouse/Zinx-WebServer/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
	len int
}

//初始话对象
func NewRequest(conn ziface.IConnection,data []byte,len int)*Request{

	res:=&Request{
		conn:conn,
		data:data,
		len:len,
	}

	return res

}



//获取链接
func(r *Request)GetConn() ziface.IConnection{

	return r.conn

}
//获取数据
func(r *Request)GetData()[]byte{

	return r.data

}
//获取数据长度
func(r *Request)GetLen()int{

	return r.len

}
