package main

import (
	"ZinxHouse/Zinx-WebServer/zinx/net"
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
	"ZinxHouse/Zinx-WebServer/MMOGame/core"
)

func OnConnAddHook(conn ziface.IConnection){
	fmt.Println("OnConnAddHook is working")
	//初始化一个玩家
	p:=core.NewPlayer(conn)

	//发送玩家id
	p.ReturnPid()

	//广播玩家位置
	p.ReturnPos()

	//
	core.WorldMagObj.AddPlayerToWorld(p)

	fmt.Println("player Id=",p.PID,"onling","player Num = ",len(core.WorldMagObj.Players))
}

func main(){
	s:=net.NewSever("MMO Game Sever")

	s.AddOnConnStart(OnConnAddHook)

	s.Sever()
}
