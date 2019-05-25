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
	fmt.Println("this is Handle")

	//_,err:=request.GetConn().GetTCPConnection().Write([]byte(" ping .......\n"))
	err:=request.GetConn().Send(1,[]byte("ping...ping...ping..."))
	if err!=nil{
		fmt.Println("Handle Write err",err)
	}

}





func main(){
	//创建服务器对象
	s:=net.NewSever("Zinx V0.5")
	//启动服务器
	s.AddRouter(&PingRouter{})

	s.Sever()

	return

}
