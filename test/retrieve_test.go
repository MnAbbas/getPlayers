package test

import (
	onefootbal "OneFootbalTest/src"
	"math"

	"testing"
)

func TestGetPlayers(t *testing.T) {

	RI := onefootbal.RetiveInfo{
		URL:      "https://vintagemonster.onefootball.com/api/teams/en/%v.json",
		MaxID:    2,
		TeamList: []string{"Apoel FC"},
	}
	ln := len(RI.FindPlayers())
	expected := 29
	if ln != expected {
		t.Errorf("expected %d , got %d", expected, ln)
	}

}

func TestGetPlayers1(t *testing.T) {

	RI := onefootbal.RetiveInfo{
		URL:      "https://vintagemonster.onefootball.com/api/teams/en/%v.json",
		MaxID:    math.MaxUint8,
		TeamList: []string{"Arsenal"},
	}
	ln := len(RI.FindPlayers())
	expected := 30
	if ln != expected {
		t.Errorf("expected %d , got %d", expected, ln)
	}

}
