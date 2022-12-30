package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Player struct {
	Name          string
	Period        string
	MatchesPlayed int
	RunsScored    int
	Avgscore      float32
}

func main() {
	var data []Player

	file, err := ioutil.ReadFile("players.json")
	if err != nil {
		fmt.Print(err)
	}

	_ = json.Unmarshal(file, &data)
	for i, _ := range data {
		data[i].MatchesPlayed = data[i].MatchesPlayed + 7
	}
	data[0].RunsScored = data[0].RunsScored + 9000
	data[1].RunsScored = data[1].RunsScored + 2000
	data[2].RunsScored = data[2].RunsScored + 1000
	slice, _ := json.MarshalIndent(data, " ", "\t")
	_ = ioutil.WriteFile("PlayerRunScoreUpdated.json", slice, 0644)

}
