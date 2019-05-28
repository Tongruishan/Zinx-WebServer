package net

import (
	"net"
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"fmt"
	"io"
	"errors"
	"ZinxHouse/Zinx-WebServer/zinx/config"
	"sync"
)

//具体的TCP链接模块
type Connection struct {

	sever ziface.ISever
	//获取链接socket
	Conn *net.TCPConn
	//获取链接Id
	ConnID uint32
	//判断是否关闭
	IsClose bool
	//绑定业务，回调函数
	//Addrouter ziface.IRouter
	MsgHandler ziface.IMsgHandler

	msgChan chan []byte
	wirterQuitChan chan bool

	ConnProperty map[string]interface{}
	ConnMutex sync.RWMutex
}

//初始化对象，相当于构造函数
func NewConnection(s ziface.ISever,conn *net.TCPConn,connId uint32,msgHandler ziface.IMsgHandler)ziface.IConnection{

	c:=&Connection{
		sever:s,
		Conn:conn,
		ConnID:connId,
		IsClose:false,
		MsgHandler:msgHandler,
		msgChan:make(chan []byte),
		wirterQuitChan:make(chan bool),
		ConnProperty:make(map[string]interface{}),

	}

	c.sever.GetConnMsg().Add(c)

	return c

}

//针对链接业务的读取方法
func(c *Connection)StartReader(){
	fmt.Println("StartReader is working ...")
	//显示信息
	defer fmt.Println("[Reader is quit ]ConnID:",c.ConnID,"Reader is quit,remote ID is ;",c.GetRemoteAddr().String())
	defer c.Stop()

	//读取
	for {
		//读取数据，所有数据都经过封装，所以需要创建封装体
		//创建封装结构体
		dp:=NewDataPack()
		//读数据需要先创建缓冲区存储数据，现在缓冲区只存储信息头
		headbuf := make([]byte,dp.GetHeadLen())
		//从数据流中读取信息，写入缓冲区中
		_,err:=io.ReadFull(c.Conn,headbuf)
		if err!=nil{
			fmt.Println("StartReader headbuf  ReadFull err",err)
			return
		}
		//将缓冲区已有的数据解封装，获取数据id和数据长度
		headdata,err:=dp.MsgUnPack(headbuf)
		if err!=nil{
			fmt.Println("StartReader MsgUnPack err",err)
			break
		}
		//将数据头的接口强转为Message结构体
		data:=headdata.(*Message)
		//为数据段开辟空间
		data.MsgData=make([]byte,data.MsgLen)
		//判断数据头长度，大于0时，读取文件内容
		if headdata.GetMsgLen()>0{
			//从流中读取数据内容
			_,err=io.ReadFull(c.Conn,data.MsgData)
			if err!=nil{
				fmt.Println("StartReader MsgData ReadFull err:",err)
				break
			}
		}

		fmt.Println("revied clinet msg: ConnId=",c.ConnID,"dataId=",data.MsgId,"datal=",string(data.GetMsgData() ))

		//创建qingqiu结构体 TODO 这个请求结构体有什么用
		req:=NewRequest(c,data)

		//传给回调函数，调用业务

		if config.GlobleConf.WorkerPoolSize>0{
			go c.MsgHandler.SendMsgToTaskSue(req)
		}else {
			go c.MsgHandler.DoMsgHandler(req)
		}



	}


}

//writer

func(c *Connection)StartWriter(){

	fmt.Println("StartWriter is working")

	defer fmt.Println("[Writer is quit ]ConnID:",c.ConnID,"remote ID is ;",c.GetRemoteAddr().String())

	for{
		select {
		case data:=<-c.msgChan:
			_,err:=c.Conn.Write(data)
			if err!=nil{
				fmt.Println("StartWriter Write err",err)
				return
			}
		case <-c.wirterQuitChan:
			return
		}
	}

}

//实现开始链接方法
func(c *Connection)Start(){
	fmt.Println("Connection Start is working ... ")

	//读取业务
	go c.StartReader()

	//写入业务
	go c.StartWriter()

	//链接之后调用钩子函数
	c.sever.CallOnConnStart(c)

}

//关闭链接方法
func(c *Connection)Stop(){

	fmt.Println("c.Stop is working ...")
	//调用钩子函数
	c.sever.CallOnConnStop(c)
	if c.IsClose==true{
		return
	}

	c.IsClose=true

	c.wirterQuitChan<-true

	err:=c.Conn.Close()
	if err!=nil{
		fmt.Println("c.close err",err)
		return
	}
	//删除链接
	c.sever.GetConnMsg().Remove(c.ConnID)

	//c.Conn.Close(c.wirterQuitChan)
	close(c.wirterQuitChan)
	close(c.msgChan)

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
func(c *Connection)Send(msgId uint32,msgData []byte)error{

	fmt.Println("Conn Send is working ...")
	//判断链接状态
	if c.IsClose==true{
		return errors.New("conn is closed!!!")
	}

	//读取数据，所有数据都经过封装，所以需要创建封装体
	//创建封装结构体
	dp:=NewDataPack()
	//将发送的内容封装
	datapack,err:=dp.MsgPack(NewMesg(msgId,msgData))
	if err!=nil{
		fmt.Println("Send MsgPack err",err)
		return err
	}

	c.msgChan<-datapack
	//
	return nil

}

func(c *Connection)SetProperty(key string,value interface{}){
	c.ConnMutex.Lock()
	defer c.ConnMutex.Unlock()

	fmt.Println("SetProperty")
	c.ConnProperty[key]=value
}

func(c *Connection)GetProperty(key string)(interface{},error){
	c.ConnMutex.RLock()
	defer c.ConnMutex.RUnlock()

	value,ok:=c.ConnProperty[key]
	if !ok{
		fmt.Println("this property is not exit")
		return nil,errors.New("this property is not exit")
	}
	return value,nil

}

func(c *Connection)RemoveProperty(key string){
	c.ConnMutex.Lock()
	defer c.ConnMutex.Unlock()

	delete(c.ConnProperty,key)

}
