package net

import (
	"net"
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
)

//具体的TCP链接模块
type Connection struct {
	//获取链接socket
	Conn *net.TCPConn
	//获取链接Id
	ConnID uint32
	//判断是否关闭
	IsClose bool
	//绑定业务，回调函数
	HandleAPI ziface.HandleFunc

}

//初始化对象，相当于构造函数
func NewConnection(conn *net.TCPConn,connId uint32,CallBackFunc ziface.HandleFunc)ziface.IConnection{

	c:=&Connection{
		Conn:conn,
		ConnID:connId,
		IsClose:false,
		HandleAPI:CallBackFunc,

	}
	return c

}

//针对链接业务的读取方法
func(c *Connection)StartReader(){
	fmt.Println("StartReader is working ...")
	//显示信息
	defer fmt.Println("ConnID:",c.ConnID,"Reader is quit,remote ID is ;",c.GetRemoteAddr().String())
	defer c.Stop()

	//读取
	for {
		//读取信息
		buf := make([]byte,512)
		n,err:=c.Conn.Read(buf)
		if err!=nil{
			fmt.Println("StartReader Read err",err)
			continue
		}
		//创建request对象
		req:=NewRequest(c,buf,n)

		//传给回调函数，调用业务
		err=c.HandleAPI(req)
		if err!=nil{
			fmt.Println("StartReader HandleAPI err",err)
			break
		}

	}

	//传给谁

}

//实现开始链接方法
func(c *Connection)Start(){
	fmt.Println("Connection Start is working ... ")

	//读取业务
	go c.StartReader()

	//TODO 写入业务

}

//关闭链接方法
func(c *Connection)Stop(){

	fmt.Println("c.Stop is working ...")

	if c.IsClose==true{
		return
	}

	c.IsClose=true

	_=c.Conn.Close()

}

//获取链接ID
func(c *Connection)GetConnId()uint32{

	return c.ConnID

}

//获取conn的原生套接字
func(c *Connection)GetTCPConnection()*net.TCPConn{
	return c.Conn
}

//获取远程Id
func(c *Connection)GetRemoteAddr()net.Addr{
	return c.Conn.RemoteAddr()
}

//发送消息
func(c *Connection)Send(data []byte,n int)error{

	fmt.Println("Conn Send is working ...")

	_,err:=c.Conn.Write(data[:n])
	if err!=nil{
		fmt.Print("Conn.Write err",err)
		return err
	}
	return nil

}
