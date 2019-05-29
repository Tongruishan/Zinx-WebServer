package core

import (
	"sync"
	"fmt"
)

const (
	AOI_MIN_X int = 85
	AOI_MAX_X int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y int = 75
	AOI_MAX_Y int = 400
	AOI_CNTS_Y int = 20
)

type WorldManager struct {
	Players map[int32]*Player
	Aoi *AOIManager
	Wmlock sync.RWMutex
}

var WorldMagObj *WorldManager

func init(){
	WorldMagObj=NewWorldManager()
}

//
func NewWorldManager()*WorldManager{

	defer fmt.Println("NewWorldManager succ!!")

	wm:=&WorldManager{
		Players:make(map[int32]*Player),
		Aoi:NewAOIManager(AOI_MIN_X,AOI_MAX_X,AOI_CNTS_X,AOI_MIN_Y,AOI_MAX_Y,AOI_CNTS_Y),
	}



	return wm

}

//
func (wm *WorldManager)AddPlayerToWorld(player *Player){



	wm.Wmlock.Lock()
	wm.Players[int32(player.PID)]=player
	wm.Wmlock.Unlock()

	//
	wm.Aoi.AddGripByPos(int(player.PID),float64(player.X),float64(player.Z))

	fmt.Println("AddPlayerToWorld succ!!")
}

//
func (wm *WorldManager)RemovePlayerToWorld(pid int32){
	wm.Wmlock.Lock()

	wm.Aoi.RemoveGripByPos(int(pid),float64(wm.Players[pid].X),float64(wm.Players[pid].Z))

	delete(wm.Players,pid)

	wm.Wmlock.Unlock()
	//

}

//
func (wm *WorldManager)GetPlayersFromWorld(pid int32)[]*Player{

	wm.Wmlock.RLock()

	plays:=make([]*Player,0)
	for _,v:=range wm.Players{

		plays=append(plays,v)
	}

	defer wm.Wmlock.RUnlock()

	return plays
}

//
func (wm *WorldManager)GetPlayerByPid(pid int32)*Player{
	wm.Wmlock.RLock()
	p:=wm.Players[pid]
	wm.Wmlock.RUnlock()



	return p

}