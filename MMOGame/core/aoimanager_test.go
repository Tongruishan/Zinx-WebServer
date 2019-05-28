package core

import (
	"testing"
	"fmt"
)

func TestAOIManager(t *testing.T) {

	aoimgr:=NewAOIManager(0,250,5,0,250,5)

	fmt.Println(aoimgr)

}

func TestAOIManagerSurround(t *testing.T) {

	aoimgr:=NewAOIManager(0,250,5,0,250,5)

	for gid,_:=range aoimgr.grips{
		gids:=aoimgr.GetSurandGidssByGid(gid)
		fmt.Println("gid : ", gid, " grids num = ", len(gids))

		gIDs:=make([]int,0,len(gids))

		for _,gid:=range gids{

			gIDs=append(gIDs,gid.GripId)

		}
		fmt.Println("grids IDs are ", gIDs)
	}

	fmt.Println("=============================")

	pids:=aoimgr.GetPidsByPos(190,50)
	fmt.Println("pids=",pids)

	//gIDs:=make([]int,0,len(aoimgr.grips))


}
