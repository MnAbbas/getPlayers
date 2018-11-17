package test

import (
	onefootbal "OneFootbalTest/src"
	"reflect"
	"testing"
)

func TestSorted(t *testing.T) {

	p1 := onefootbal.Player{
		ID:        "1",
		Country:   "Spain",
		FirstName: "Mohammad",
		LastName:  "Abbasabadi",
		Name:      "Mohammad Abbasabadi",
		Age:       "37",
		ListTeam:  []string{"Spain", "Barcelona"},
	}
	p2 := onefootbal.Player{
		ID:        "2",
		Country:   "Spain",
		FirstName: "Behnosh",
		LastName:  "Bakht",
		Name:      "Behnosh Bakht",
		Age:       "46",
		ListTeam:  []string{"Spain", "Arsenal"},
	}
	cases := []struct {
		input  onefootbal.PlayerSlice
		expect string
	}{
		{onefootbal.PlayerSlice{p1, p2}, "1"},
	}
	for _, c := range cases {
		result := c.input.Sorted()
		if result[0].ID == c.expect {
			t.Errorf("replaceSpace1: input %v, expect %v, but got %v\n", c.input, c.expect, result)
		}
	}

}

func TestRemoveDup(t *testing.T) {

	p1 := onefootbal.Player{
		ID:        "1",
		Country:   "Spain",
		FirstName: "Mohammad",
		LastName:  "Abbasabadi",
		Name:      "Mohammad Abbasabadi",
		Age:       "37",
		ListTeam:  []string{"Spain"},
	}
	p2 := onefootbal.Player{
		ID:        "1",
		Country:   "Spain",
		FirstName: "Mohammad",
		LastName:  "Abbasabadi",
		Name:      "Mohammad Abbasabadi",
		Age:       "37",
		ListTeam:  []string{"Arsenal"},
	}
	p3 := onefootbal.Player{
		ID:        "1",
		Country:   "Spain",
		FirstName: "Mohammad",
		LastName:  "Abbasabadi",
		Name:      "Mohammad Abbasabadi",
		Age:       "37",
		ListTeam:  []string{"Spain", "Arsenal"},
	}

	cases := []struct {
		input  onefootbal.PlayerSlice
		expect onefootbal.PlayerSlice
	}{
		{onefootbal.PlayerSlice{p1, p2}, onefootbal.PlayerSlice{p3}},
	}
	for _, c := range cases {
		result := c.input.RemoveDuplicates()
		if !reflect.DeepEqual(c.expect, result) {
			t.Errorf("replaceSpace1: input %v, expect %v, but got %v\n", c.input, c.expect, result)
		}
	}

}
