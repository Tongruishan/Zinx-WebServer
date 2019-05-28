package core

import (
	"sync"
	"fmt"
)

//格子类
type Grip struct {
	GripId int
	MinX int
	MaxX int
	MinY int
	MaxY int
	PlayerIds map[int]interface{}
	PlayerIdLock sync.RWMutex
}

//初始化格子
func NewGrip(gripId int,minx int,maxx int,miny int,maxy int)*Grip{

	return &Grip{
		GripId:gripId,
		MinX:minx,
		MaxX:maxx,
		MinY:miny,
		MaxY:maxy,
		PlayerIds:make(map[int]interface{}),
	}

}

//玩家进入格子
func(g *Grip)Add(playerId int,player interface{}){
	g.PlayerIdLock.Lock()
	defer g.PlayerIdLock.Unlock()

	g.PlayerIds[playerId]=player

}

//玩家离开格子
func(g *Grip)Remove(playerId int){
	g.PlayerIdLock.Lock()
	defer g.PlayerIdLock.Unlock()

	delete(g.PlayerIds,playerId)

}

//查看格子内所有玩家
func(g *Grip)GetAllPlayers()(playIds []int){
	g.PlayerIdLock.RLock()
	defer g.PlayerIdLock.RUnlock()

	for playId,_:=range g.PlayerIds{
		//只需要key，不需要value的，所以将map内容存到value里
		playIds=append(playIds,playId)
	}

	return

}

//打印信息
func(g *Grip)String()string{
	return fmt.Sprintf("PlayerId:%d,MinX:%d,MaxX:%d,MinY:%d,MaxY:%d,PlayerIds:%v\n",
											g.GripId,g.MinX,g.MaxX,g.MinY,g.MaxY,g.PlayerIds)
}

