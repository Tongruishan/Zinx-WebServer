package main

import (
	"ZinxHouse/Zinx-WebServer/zinx/net"
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
)

//import "ZinxHouse/Zinx-WebServer/zinx/net"
//

type PingRouter struct {
	net.BaseRouter
}

func(p *PingRouter)PreHandle(request ziface.IRequest){
	fmt.Println("this is PreHandle")
	_,err:=request.GetConn().GetTCPConnection().Write([]byte("before ping .......\n"))
	if err!=nil{
		fmt.Println("PreHandle Write err",err)
	}
}

func(p *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("this is Handle")
	_,err:=request.GetConn().GetTCPConnection().Write([]byte(" ping .......\n"))
	if err!=nil{
		fmt.Println("Handle Write err",err)
	}

}

func(p *PingRouter)PostHandle(request ziface.IRequest){
	fmt.Println("this is PostHandle")
	_,err:=request.GetConn().GetTCPConnection().Write([]byte("After ping .......\n"))
	if err!=nil{
		fmt.Println("AfterHandle Write err",err)
	}

}



func main(){
	//创建服务器对象
	s:=net.NewSever("Zinx V1.0")
	//启动服务器
	s.AddRouter(&PingRouter{})

	s.Sever()

	return

}
