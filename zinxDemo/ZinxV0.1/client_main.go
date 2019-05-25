package main

import (
	"net"
	"fmt"
	"time"
	net2 "ZinxHouse/Zinx-WebServer/zinx/net"
	"io"
)

func main(){

	fmt.Println("client connect ....")

	time.Sleep(time.Second)

	conn,err:=net.Dial("tcp","127.0.0.1:8999")
	if err!=nil{
		fmt.Println("Dial err:",err)
		return
	}

	for{

		//向客户端发送的消息
		dp:=net2.NewDataPack()
		data,err:=dp.MsgPack(net2.NewMesg(0,[]byte("hello zinx")))
		if err!=nil{
			fmt.Println("MsgPack err",err)
			return
		}
		_,err=conn.Write(data)
		if err!=nil{
			fmt.Println("Write err",err)
			return
		}

		//接受客户端发送的消息

		databuf:=make([]byte,dp.GetHeadLen())
		_,err=io.ReadFull(conn,databuf)
		if err!=nil{
			fmt.Println("ReadFull err",err)
			return
		}

		dataHead,err:=dp.MsgUnPack(databuf)
		if err!=nil{
			fmt.Println("MsgUnPack err",err)
			return
		}

		if dataHead.GetMsgLen()>0{
			msg:=dataHead.(*net2.Message)
			msg.MsgData=make([]byte,msg.MsgLen)

			_,err=io.ReadFull(conn,msg.MsgData)
			if err!=nil{
				fmt.Println("MsgUnPack ReadFull err",err)
				return
			}

			fmt.Println("sever call msgId=",msg.MsgId,"msgLen=",msg.MsgLen,"msgData=",string(msg.MsgData))
		}

		time.Sleep(time.Second)
	}
}