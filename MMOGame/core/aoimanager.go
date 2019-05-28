package core

import "fmt"

//import "openbilibili/go-common-master/app/service/main/member/api"

type AOIManager struct {
	MinX int
	MAXX int
	ContX int
	MinY int
	MAXY int
	Conty int
	grips map[int]*Grip


}

//获取地图宽度，x轴方格个数
func (am *AOIManager)GripWeith()int  {

	return (am.MAXX-am.MinX)/am.ContX

}

//获取地图高度，y轴方格个数
func (am *AOIManager)GripHeith()int  {

	return (am.MAXY-am.MinY)/am.Conty

}

//初始化aoi管理器
func NewAOIManager(minx,maxx,contx,miny,maxy,conty int)*AOIManager{
	 s:=&AOIManager{
		MinX:minx,
		MAXX:maxx,
		ContX:contx,
		MinY:miny,
		MAXY:maxy,
		Conty:conty,
		grips:make(map[int]*Grip),
	}

	//初始化每个方格
	for y:=0;y<conty;y++{
		for x:=0;x<contx;x++{

			gip:=y*contx+x

			s.grips[gip]= NewGrip(gip,
				s.MinX+x*s.GripWeith(),
				s.MinX+(x+1)*s.GripWeith(),
				s.MinY+y*s.GripHeith(),
				s.MinY+(y+1)*s.GripHeith(),
				)

		}
	}
	return s
}

//打印信息
func(am *AOIManager)String()string{

	s:=fmt.Sprintf("AOIManger:\n Minx:%d,Maxx:%d,ContX:%d,MinY%d,MaxY:%d,ContY:%d Grods in manager",
		am.MinX,am.MAXX,am.ContX,am.MinY,am.MAXY,am.Conty)

	for _,Gid :=range am.grips{
		s+=fmt.Sprintln(Gid)
	}

	return s

}

//玩家进入方格
func(am *AOIManager)AddPidToGrip(pid,gid int){
	am.grips[gid].Add(pid,nil)

}

//玩家离开方格
func(am *AOIManager)RemovePidFromGrip(pid,gid int){
	am.grips[gid].Remove(pid)
}

//获取当前方格内所有玩家
func(am *AOIManager)GetPidByGid(gid int)(pids []int){

	pids=am.grips[gid].GetAllPlayers()

	return
}


//通过当前方格id获取周围所有方格id
func(am *AOIManager)GetSurandGidssByGid(gid int)(gids []*Grip){

	_,ok:=am.grips[gid]
	if !ok{
		return
	}

	gids=append(gids,am.grips[gid])

	idX:=gid%am.ContX

	if idX>0{
		gids=append(gids,am.grips[gid-1])
	}
	if idX<am.ContX-1{
		gids=append(gids,am.grips[gid+1])
	}

	gidsX:=make([]int,0,len(gids))

	for _,V:=range gids{
		gidsX=append(gidsX,V.GripId)
	}

	for _,gid:=range gidsX{

		idY:=gid/am.ContX

		if idY>0{
			gids=append(gids,am.grips[gid-am.ContX])
		}
		if idY<am.Conty-1{
			gids=append(gids,am.grips[gid+am.ContX])
		}

	}

	return

}

//通过坐标获取所在方格id
func(am *AOIManager)GetGidByPos(x,y float64)int{

	if x<0||int(x)>am.ContX{
		return -1
	}
	if y<0||int(y)>am.Conty{
		return -1
	}

	px:=(int(x)-am.MinX)/am.GripWeith()

	py:=(int(y)-am.MinY)/am.GripHeith()

	gid:=py*am.ContX+px

	return gid
}

//通过坐标获取周围所有玩家
func(am *AOIManager)GetPidsByPos(x,y float64)(pids []int){

	gid:=am.GetGidByPos(x,y)
	gids:=am.GetSurandGidssByGid(gid)

	for _,gid:=range gids{
		pids1:=gid.GetAllPlayers()
		pids=append(pids,pids1...)
	}

	return

}


//通过坐标，将玩家添加到方格
func(am *AOIManager)AddGripByPos(pid int,x,y float64){

	gid:=am.GetGidByPos(x,y)

	grid:=am.grips[gid]

	grid.Add(pid,nil)

}

//通过坐标将玩家从方格删除
func(am *AOIManager)RemoveGripByPos(pid int,x,y float64){
	gid:=am.GetGidByPos(x,y)

	grid:=am.grips[gid]

	grid.Remove(pid)
}



