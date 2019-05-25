package net

import (
	"testing"
	"fmt"
	"net"
	"io"
)

func TestDataPack(t *testing.T){
	fmt.Println(" test datapack ......")

	//服务器
	listener,err:=net.Listen("tcp","127.0.0.1:7878")
	if err!=nil{
		fmt.Println("Listen err",err)
		return
	}

	go func() {
		for{
			conn,err:=listener.Accept()
			if err!=nil{
				fmt.Println("Accept err",err)
				//break
			}

			//读写业务
			go func(conn *net.Conn) {

				for{
					//拆包，创建拆包对象
					dp:=NewDataPack()

					headData:=make([]byte,dp.GetHeadLen())

					_,err:=io.ReadFull(*conn,headData)
					if err!=nil{
						fmt.Println("ReadFull headData err",err)
						return
					}

					msgHead,err:=dp.MsgUnPack(headData)
					if err!=nil{
						fmt.Println("MsgUnPack err",err)
						return
					}

					//判断解包成功
					if msgHead.GetMsgLen()>0{
						//msgHead强转未msg
						msg:=msgHead.(*Message)
						//未data开辟空间
						msg.MsgData=make([]byte,msgHead.GetMsgLen())

						//向打他中写入数据据
						_,err:=io.ReadFull(*conn,msg.MsgData)
						if err!=nil{
							fmt.Println("MsgData ReadFull err",err)
							return
						}

						fmt.Println("revec MsgId:",msg.MsgId,"MsgLen:",msg.MsgLen,"MsgData:",string( msg.MsgData))
					}
				}
			}(&conn)
		}
	}()


	//客户端
	cliConn,err:=net.Dial("tcp","127.0.0.1:7878")
	if err!=nil{
		fmt.Println("cliConn Dial err",err)
		return
	}

	//封装数据
	dp2:=NewDataPack()

	msg1:=&Message{
		MsgId:1,
		MsgLen:7,
		MsgData:[]byte{'b','e','i','j','i','n','g'},
	}
	data1,err:=dp2.MsgPack(msg1)
	if err!=nil{
		fmt.Println("cliConn MsgPack err",err)
		return
	}

	msg2:=&Message{
		MsgId:1,
		MsgLen:4,
		MsgData:[]byte{'t','o','n','g'},
	}
	data2,err:=dp2.MsgPack(msg2)
	if err!=nil{
		fmt.Println("cliConn MsgPack err",err)
		return
	}

	data1=append(data1,data2...)


	cliConn.Write(data1)

	select {

	}

}
