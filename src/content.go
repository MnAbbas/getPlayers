package onefootbal

import (
	"runtime"
	"sync"
)

const (
	url = "https://vintagemonster.onefootball.com/api/teams/en/%v.json"
	//MaxUint shows maximum uint id could be one of ^uint(0) , ^uint16(0) , ^uint32(0)
	MaxUint = ^uint8(0)
)

var wg sync.WaitGroup

//RetiveInfo is acted as facade layer
type RetiveInfo struct {
	URL       string
	MaxID     uint
	TeamList  []string
	RemoveDup bool
	Sorting   bool
}

//FindPlayers do all stuff together
func (ri RetiveInfo) FindPlayers() PlayerSlice {
	runtime.GOMAXPROCS(runtime.NumCPU())

	acceptedTeams := map[string]bool{}

	for _, v := range ri.TeamList {
		acceptedTeams[v] = true
	}

	var p PlayerSlice
	var index uint
	for index = 1; index < ri.MaxID; index++ {
		if len(acceptedTeams) == 0 {
			break
		}
		wg.Add(1)
		teamname, players := ri.Getcontent(index)
		if acceptedTeams[teamname] {
			delete(acceptedTeams, teamname)
			p = append(p, players...)
		}
	}
	wg.Wait()
	if ri.RemoveDup && ri.Sorting {
		return p.SortAndConcat()
	}
	if ri.RemoveDup && !ri.Sorting {
		return p.RemoveDuplicates()
	}
	if !ri.RemoveDup && ri.Sorting {
		return p.Sorted()
	}
	return p
}

//Run is the main function here
func Run() {
	acceptedT := []string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea", "Barcelona", "Real Madrid", "FC Bayern Munich"}
	ri := RetiveInfo{
		URL:       url,
		MaxID:     uint(MaxUint),
		TeamList:  acceptedT,
		RemoveDup: true,
		Sorting:   true}
	players := ri.FindPlayers()
	players.FormatOutPut()
}
