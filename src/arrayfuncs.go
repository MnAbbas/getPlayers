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

//Sorted is for sorting using heap
func (p PlayerSlice) Sorted() PlayerSlice {
	sort.Sort(p)
	return p
}

//SortAndConcat is for Sorting and Removing Duplicate and concat them
func (p PlayerSlice) SortAndConcat() PlayerSlice {
	newarr := p.RemoveDuplicates()
	sort.Sort(newarr)
	return newarr
}

//FormatOutPut is for formating output
func (p PlayerSlice) FormatOutPut() {

	for index := 0; index < len(p); index++ {
		elemnt := p[index]
		fmt.Printf("%d. %v %v; %v; %v\n", index+1, elemnt.FirstName, elemnt.LastName, elemnt.Age, strings.Join(elemnt.ListTeam[:], ","))

	}
}

//RemoveDuplicates is for remove Duplicate itemss and concat them
func (p PlayerSlice) RemoveDuplicates() PlayerSlice {
	refhappend := map[string]Player{}
	result := []Player{}
	for _, v := range p {
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
