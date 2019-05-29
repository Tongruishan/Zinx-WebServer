package core

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"sync"
	"math/rand"
	"github.com/golang/protobuf/proto"
	//"github.com/astaxie/beego"
	"fmt"
	"ZinxHouse/Zinx-WebServer/MMOGame/pb"
)

//创建玩家对象
type Player struct {
	PID uint32
	Conn ziface.IConnection
	X float32
	Y float32
	Z float32
	V float32
}

//玩家id累加器
var Pid uint32=1
var PidLock sync.Mutex

//初始化玩家
func NewPlayer(conn ziface.IConnection)*Player{
	PidLock.Lock()
	id:=Pid
	Pid++
	PidLock.Unlock()

	p:=&Player{
		PID:id,
		Conn:conn,
		X:float32(100+rand.Intn(10)),
		Y:0,
		Z:float32(100+rand.Intn(10)),
		V:0,
	}

	return p

}

//玩家发送消息模块
func(p *Player)SendMsg(msgId uint32,ProtoMessage proto.Message)error{

	data,err:=proto.Marshal(ProtoMessage)
	if err!=nil{
		fmt.Println("Marshal err",err)
		return err
	}


	err=p.Conn.Send(msgId,data)
	if err!=nil{
		fmt.Println("Send err",err)
		return err
	}

	return nil
}

//发送玩家id消息
func(p *Player)ReturnPid(){

	pid:=&pb.SyncPid{
		Pid:1,
	}

	p.SendMsg(1,pid)
}

//发送玩家位置消息
func(p *Player)ReturnPos(){

	pos:=&pb.BroadCast{
		Pid:int32(p.PID),
		Tp:2,
		Data:&pb.BroadCast_P{
			&pb.Position{
				X:p.X,
				Y:p.Y,
				Z:p.Z,
				V:p.V,
			},
		},
	}

	p.SendMsg(200,pos)
}
