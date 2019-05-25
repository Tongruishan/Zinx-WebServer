package main

import (
	"ZinxHouse/Zinx-WebServer/zinx/net"
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
)


type PingRouter struct {
	net.BaseRouter
}

func(p *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("this is PingRouter Handle")

	//_,err:=request.GetConn().GetTCPConnection().Write([]byte(" ping .......\n"))
	err:=request.GetConn().Send(200,[]byte("ping...ping...ping..."))
	if err!=nil{
		fmt.Println("Handle Write err",err)
	}

}


type DongRouter struct {
	net.BaseRouter
}

func(p *DongRouter)Handle(request ziface.IRequest){
	fmt.Println("this is DongRouter Handle")

	//_,err:=request.GetConn().GetTCPConnection().Write([]byte(" ping .......\n"))
	err:=request.GetConn().Send(201,[]byte("dong...dong...dong..."))
	if err!=nil{
		fmt.Println("Handle Write err",err)
	}

}


func main(){
	//创建服务器对象
	s:=net.NewSever("Zinx V0.5")
	//启动服务器
	s.AddRouter(1,&PingRouter{})
	s.AddRouter(2,&DongRouter{})

	s.Sever()

	return

}
