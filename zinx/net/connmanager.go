package net

import (
	"ZinxHouse/Zinx-WebServer/zinx/ziface"
	"sync"
	"fmt"
	"errors"
)

type ConnManager struct {
	Connections map[uint32]ziface.IConnection
	//对map和集合操作一定要枷锁
	RWMutex sync.RWMutex
}

func NewConnManager()ziface.IConnManager{

	return &ConnManager{
		Connections:make(map[uint32]ziface.IConnection),
	}

}

func(cm *ConnManager) Add(conn ziface.IConnection){
	cm.RWMutex.Lock()
	defer cm.RWMutex.Unlock()

	cm.Connections[conn.GetConnId()]=conn

	fmt.Println("Connd=",conn.GetConnId(),"add succ!!!")
}

func(cm *ConnManager) Remove(connId uint32){
	cm.RWMutex.Lock()
	defer cm.RWMutex.Unlock()

	delete(cm.Connections,connId)
	fmt.Println("Connd=",connId,"remove succ!!!")

}
func(cm *ConnManager) Get(connId uint32)(ziface.IConnection,error){
	cm.RWMutex.RLock()
	defer cm.RWMutex.RUnlock()

	conn,ok:=cm.Connections[connId]
	if ok{
		return conn,nil
	}else {
		return nil,errors.New("NOT FOUND ")
	}

}
func(cm *ConnManager) Len()int{
	cm.RWMutex.RLock()
	defer cm.RWMutex.RUnlock()

	len:=len(cm.Connections)

	return len


}
func(cm *ConnManager) ClearConn(){
	cm.RWMutex.Lock()
	defer cm.RWMutex.Unlock()

	for connId,conn:=range cm.Connections{
		conn.Stop()
		delete(cm.Connections,connId)
	}

	fmt.Println("clear conn over ")

}

