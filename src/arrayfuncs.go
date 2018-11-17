package onefootbal

import (
	"fmt"
	"sort"
	"strings"
)

//PlayerSlice is for creating a heap
type PlayerSlice []Player

func (p PlayerSlice) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p PlayerSlice) Len() int           { return len(p) }
func (p PlayerSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (s PlayerSlice) Sorted() PlayerSlice {
	sort.Sort(s)
	return s
}

func (s PlayerSlice) SortAndConcat() PlayerSlice {
	newarr := s.RemoveDuplicates()
	sort.Sort(newarr)
	return newarr
}

func (players PlayerSlice) FormatOutPut() {

	for index := 0; index < len(players); index++ {
		elemnt := players[index]
		fmt.Printf("%d. %v %v; %v; %v\n", index+1, elemnt.FirstName, elemnt.LastName, elemnt.Age, strings.Join(elemnt.ListTeam[:], ","))

	}
}

func (players PlayerSlice) RemoveDuplicates() PlayerSlice {
	refhappend := map[string]Player{}
	result := []Player{}
	for _, v := range players {
		if exist := refhappend[v.ID]; exist.ID != "" {
			teamname := []string{exist.ListTeam[0], v.ListTeam[0]}
			v.ListTeam = teamname
			refhappend[v.ID] = v
		} else {
			refhappend[v.ID] = v
		}
	}
	for _, v := range refhappend {
		result = append(result, v)
	}
	return result
}
