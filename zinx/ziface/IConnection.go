package ziface

import "net"

/*
	抽象层
*/

type IConnection interface {
	//
	Start()

	//
	Stop()

	//获取链接ID
	GetConnId()uint32
	
	//获取conn的原生套接字
	GetTCPConnection()*net.TCPConn

	//获取远程Id
	GetRemoteAddr()net.Addr

	//发送消息
	Send(data []byte)error


}

//业务处理的抽象方法
type HandleFunc func(*net.TCPConn,[]byte,int) error
