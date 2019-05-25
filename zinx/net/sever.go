package net

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"net"
	"fmt"
	"ZinxHouse/Zinx-WebServer/zinx/config"
)

//创建对象
type Sever struct{
	IpVersion string
	Ip string
	Port int
	Name string
	Addrouter ziface.IRouter
}

//对象初始化
func NewSever(name string) ziface.ISever {

	S:= &Sever{
		Name:config.GlobleConf.Name,
		IpVersion:"tcp4",
		Ip:config.GlobleConf.Host,
		Port:config.GlobleConf.Port,
		Addrouter:nil,
	}
	return S

}


////回调函数
//func CallBackBusi(request ziface.IRequest)error{
//	fmt.Println("CallBackBusi is working")
//
//	conn:=request.GetConn().GetTCPConnection()
//	data:=request.GetData()
//	n:=request.GetLen()
//
//
//	_,err:=conn.Write(data[:n])
//	if err!=nil{
//		fmt.Println("CallBackBusi Write err",err)
//		return err
//	}
//
//	return nil
//
//}en

//对象方法
//停止服务
func(this *Sever)Start(){
	//链接服务器
	addr,err:=net.ResolveTCPAddr("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err!=nil{
		fmt.Println("ResolveTCPAddr err:",err)
		return
	}
	//获取监听器
	linstener,err:=net.ListenTCP(this.IpVersion,addr)
	if err!=nil{
		fmt.Println("ListenTCP err:",err)
		return
	}

	var cid uint32
	cid=0

	go func() {
		for {
			//开始监听
			conn,err:=linstener.AcceptTCP()
			if err!=nil{
				fmt.Println("Accept err:",err)
				continue
			}
			//调用链接模块
			//delConn:=NewConnection(conn,cid,CallBackBusi)
			delConn:=NewConnection(conn,cid,this.Addrouter)
			cid++


			//链接模块的开始链接方法
			go delConn.Start()

		}
	}()

}

//开始服务
func(this *Sever)Stop(){

}

//服务
func(this *Sever)Sever(){
	//调用start
	this.Start()

	//TODO 其他的事,防止主go程推出
	select {

	}
}

//路由，将sever对象自己的属性和对象的路由建立链接，非常中国要
func(this *Sever)AddRouter(router ziface.IRouter){
	this.Addrouter=router
}