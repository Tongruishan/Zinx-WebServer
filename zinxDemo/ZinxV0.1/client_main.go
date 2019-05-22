package main

import (
	"net"
	"fmt"
	"time"
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

		_,err=conn.Write([]byte("hello zinx ...."))
		if err!=nil{
			fmt.Println("Write err:",err)
			return
		}

		buf:=make([]byte,512)
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("Read err:",err)
			return
		}
		fmt.Printf("sever call back %s,%d\n",buf,n)

		time.Sleep(time.Second)
	}
}