package onefootbal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//Message is a message render from get url
type Message struct {
	Status  string
	Code    int
	Data    Data
	Message string
}

//Data structure of inneral object from get url
type Data struct {
	Team Team
}

//Team struc for Team data from get url
type Team struct {
	ID      int
	Name    string
	Players []Player
}

//Player struct for represent player informtion
type Player struct {
	ID        string
	Country   string
	FirstName string
	LastName  string
	Name      string
	Age       string
	ListTeam  []string
}

//getcontent to retrieve data from url
func (ri RetiveInfo) Getcontent(teamid uint) (string, []Player) {
	// json data
	url := fmt.Sprintf(ri.URL, teamid)

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(url)
	if err != nil {
		wg.Done()
		return "", []Player{}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		wg.Done()
		return "", []Player{}
	}
	defer res.Body.Close()
	var data Message
	json.Unmarshal(body, &data)
	for i := 0; i < len(data.Data.Team.Players); i++ {
		var teamnames []string
		teamnames = []string{data.Data.Team.Name}
		data.Data.Team.Players[i].ListTeam = teamnames
	}
	wg.Done()
	return data.Data.Team.Name, data.Data.Team.Players
}
