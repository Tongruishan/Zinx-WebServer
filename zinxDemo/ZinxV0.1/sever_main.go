package main

//import "ZinxHouse/Zinx-WebServer/zinx/net"
//

import "ZinxHouse/Zinx-WebServer/zinx/net"

func main(){
	//创建服务器对象
	s:=net.NewSever("Zinx V1.0")
	//启动服务器
	s.Sever()
}
