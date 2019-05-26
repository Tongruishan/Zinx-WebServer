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
	MsgHandler ziface.IMsgHandler
	ConnMag ziface.IConnManager
	//创建2个钩子函数置身
	OnConnStart func(conn ziface.IConnection)
	OnConnStop func(conn ziface.IConnection)
}

//对象初始化
func NewSever(name string) ziface.ISever {

	S:= &Sever{
		Name:config.GlobleConf.Name,
		IpVersion:"tcp4",
		Ip:config.GlobleConf.Host,
		Port:config.GlobleConf.Port,
		MsgHandler:NewMsgHandler(),
		ConnMag:NewConnManager(),
	}
	return S

}


//对象方法
//停止服务
func(this *Sever)Start(){

	fmt.Println("Sever is start ......")

	this.MsgHandler.StartWorkPool()
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
			//判断链接数量
			if this.ConnMag.Len()>int(config.GlobleConf.MaxConn){
				fmt.Println("too many conn,this conn is closed")
				conn.Close()
				continue
			}

			delConn:=NewConnection(this,conn,cid,this.MsgHandler)
			cid++

			//链接模块的开始链接方法
			go delConn.Start()

		}
	}()

}

//开始服务
func(this *Sever)Stop(){

	fmt.Println("Sever is stop")

	//清除所有链接
	this.ConnMag.ClearConn()

}

//服务
func(this *Sever)Sever(){

	fmt.Println("Sever is severing")
	//调用start
	this.Start()

	//防止主go程推出
	select {

	}
}

//路由，将sever对象自己的属性和对象的路由建立链接，非常中国要
func(this *Sever)AddRouter(msgId uint32,router ziface.IRouter){
	this.MsgHandler.AddMsgRouter(msgId,router)
	fmt.Println("msgId=",msgId,"router",router,"has been apeend")


}

func(this *Sever)GetConnMsg()ziface.IConnManager{

	return this.ConnMag

}

//zhuce函数

func(this *Sever)AddOnConnStart(hookF func(conn ziface.IConnection)){

	this.OnConnStart=hookF

}

func(this *Sever)AddOnConnStop(hookF func(conn ziface.IConnection)){
	this.OnConnStop=hookF
}


//调用函数
func(this *Sever)CallOnConnStart(conn ziface.IConnection){

	if this.OnConnStart!=nil{
		fmt.Println("----> CallOnConnStart ")
		this.OnConnStart(conn)
	}

}

func(this *Sever)CallOnConnStop(conn ziface.IConnection){
	if this.OnConnStop!=nil{
		fmt.Println("----> CallOnConnStop ")
		this.OnConnStop(conn)
	}

}