package net

import (
	"ZinxHouse/zinx/ziface"
	"net"
	"fmt"
)

//创建对象
type Sever struct{
	IpVersion string
	Ip string
	Port int
	Name string
}

//对象初始化
func NewSever(name string)  ziface.ISever {

	S:= Sever{
		Name:name,
		IpVersion:"tcp4",
		Ip:"0.0.0.0",
		Port:8999,
	}
	return &S

}

//对象方法
//停止服务
func(this *Sever)Stop(){

	addr,err:=net.ResolveTCPAddr("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err!=nil{
		fmt.Println("ResolveTCPAddr err:",err)
		return
	}

	linstener,err:=net.ListenTCP(this.IpVersion,addr)
	if err!=nil{
		fmt.Println("ListenTCP err:",err)
		return
	}

	go func() {
		for {
			conn,err:=linstener.Accept()
			if err!=nil{
				fmt.Println("Accept err:",err)
				continue
			}

			go func() {



				for {

					//buf=[]byte{}
					//buf = buf[:0]
					buf:=make([]byte,512)
					n,err:=conn.Read(buf)
					if err!=nil{
						fmt.Println("Read err:",err)
						break
					}

					//相当于写日志
					fmt.Printf("recv client buf %s, cnt = %d\n", buf, n)

					_,err=conn.Write(buf[:n])
					if err!=nil{
						fmt.Println("Write err:",err)
						continue
					}
				}

			}()

		}
	}()




}


//开始服务
func(this *Sever)Start(){

}

//服务
func(this *Sever)Sever(){
	//调用start
	this.Stop()

	//TODO 其他的事
	select {

	}
}