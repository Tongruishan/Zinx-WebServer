package net

import (
	"net"
	"ZinxHouse/zinx/ziface"
)

//具体的TCP链接模块
type Connection struct {

	Conn *net.TCPConn

	ConnID uint32

	IsClose bool

	HandleAPI ziface.HandleFunc

}


//实现方法

func(c *Connection)Start(){

}

//
func(c *Connection)Stop(){

}

//获取链接ID
func(c *Connection)GetConnId()uint32{

	return 0

}

//获取conn的原生套接字
func(c *Connection)GetTCPConnection()*net.TCPConn{
	return nil
}

//获取远程Id
func(c *Connection)GetRemoteAddr()net.Addr{
	return nil
}

//发送消息
func(c *Connection)Send(data []byte)error{
	return nil

}
