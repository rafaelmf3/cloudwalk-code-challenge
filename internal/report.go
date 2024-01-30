package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	file "github.com/rafaelmf3/parse/pkg/file"
	"github.com/rafaelmf3/parse/pkg/player"
)

type GameReport struct {
	TotalKills  int32
	PlayerKills []player.Player
}

func CreateJsonGameFile(games map[int]*Game) {
	gamesReport := make(map[string]Game)
	for _, v := range games {
		gamesReport["game-"+strconv.Itoa(v.ID)] = Game{
			ID:          v.ID,
			TotalKills:  v.TotalKills,
			PlayerKills: v.PlayerKills,
			Kills:       v.Kills,
		}
	}

	data, err := json.MarshalIndent(gamesReport, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	file.CreateFile("game_json_file.json", data)
}

func CreateGameReport() {
	f, _ := os.ReadFile("game_json_file.json")
	var games map[string]*Game
	json.Unmarshal(f, &games)
	gamesReport := make(map[string]GameReport)

	fmt.Printf("%+v", games)
	for _, v := range games {
		pk := player.RankPlayers(v.PlayerKills)
		gamesReport["game_"+strconv.Itoa(v.ID)] = GameReport{
			TotalKills:  v.TotalKills,
			PlayerKills: pk,
		}
	}

	data, err := json.MarshalIndent(gamesReport, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	file.CreateFile("games_report.json", data)
}

func CreateJsonDeathCauseReport() {
	f, _ := os.ReadFile("game_json_file.json")
	var games map[string]*Game
	json.Unmarshal(f, &games)
	report := make(map[string]DeathCauseReport)
	for _, v := range games {
		killsByMean := make(map[string]int)
		for _, k := range v.Kills {
			for _, d := range k {
				killsByMean[d] += 1
			}
		}
		if len(killsByMean) > 0 {
			report["game_"+strconv.Itoa(v.ID)] = DeathCauseReport{KillsByMean: killsByMean}
		}
	}

	data, err := json.MarshalIndent(report, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	file.CreateFile("death_cause_report.json", data)
}
