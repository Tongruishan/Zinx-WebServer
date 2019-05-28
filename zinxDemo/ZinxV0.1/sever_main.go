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
	fmt.Println(request.GetMsg().GetMsgData())

	err:=request.GetConn().Send(201,[]byte("dong...dong...dong..."))
	if err!=nil{
		fmt.Println("Handle Write err",err)
	}

}

func DoConnBeginn(conn ziface.IConnection){

	fmt.Println("=====>DoConnBeginn")
	err:=conn.Send(301,[]byte("This is DoConnBeginn,"))
	if err!=nil{
		fmt.Println("DoConnBeginn Send err ")
	}
	fmt.Println("------->SetProperty succ1")
	conn.SetProperty("name","tongxiaotong")
	conn.SetProperty("age",29)
	conn.SetProperty("sex","man")
	fmt.Println("------->SetProperty succ2")

}

func DoConnStop(conn ziface.IConnection){
	fmt.Println("=====>DoConnStop")
	err:=conn.Send(302,[]byte("This is DoConnStop,"))
	if err!=nil{
		fmt.Println("DoConnBeginn Send err ")
	}
	if name,ok:=conn.GetProperty("name");ok==nil{
		fmt.Println("name=",name)
	}
	if age,ok:=conn.GetProperty("age");ok==nil{
		fmt.Println("age=",age)
	}
	if sex,ok:=conn.GetProperty("sex");ok==nil{
		fmt.Println("sex=",sex)
	}

	fmt.Println("------->GetProperty succ")


}


func main(){
	//创建服务器对象
	s:=net.NewSever("Zinx V0.5")
	s.AddOnConnStart(DoConnBeginn)
	s.AddOnConnStop(DoConnStop)
	//启动服务器
	s.AddRouter(1,&PingRouter{})
	s.AddRouter(2,&DongRouter{})

	s.Sever()

	return

}
