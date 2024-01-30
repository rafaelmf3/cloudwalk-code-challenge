package internal

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Game struct {
	ID          int                 `json:"id,omitempty"`
	TotalKills  int32               `json:"total_kills"`
	PlayerKills map[string]int      `json:"player_kills"`
	Kills       map[string][]string `json:"kills,omitempty"`
}

type DeathCauseReport struct {
	KillsByMean map[string]int `json:"kills_by_means,omitempty"`
}

func NewGame(gamenumber int) *Game {
	return &Game{
		ID:          gamenumber,
		TotalKills:  0,
		PlayerKills: make(map[string]int, 0),
		Kills:       make(map[string][]string),
	}
}

func (g *Game) Kill(killer, killed, deathCause string) {
	g.TotalKills++
	g.Kills[killer] = append(g.Kills[killer], deathCause)
	if killer == "<world>" {
		g.PlayerKills[killed]--
	} else {
		g.PlayerKills[killer]++
	}
}

func ScanGames(file *os.File) map[int]*Game {

	scanner := bufio.NewScanner(file)
	game := 0
	games := make(map[int]*Game)

	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		lineText := string(lineBytes)

		killedPlayer := []string{}
		killerPlayer := []string{}

		killed, killer := false, true

		if bytes.Contains(lineBytes, []byte("InitGame")) {
			game++
			games[game] = NewGame(game)
		}

		if bytes.Contains(lineBytes, []byte("Kill:")) {
			payload := strings.Split(strings.TrimLeft(strings.Split(lineText, ":")[3], " "), " ")

			for _, v := range payload {
				if v == "by" {
					killed = false
				}

				if killed {
					killedPlayer = append(killedPlayer, v)
				}

				if v == "killed" {
					killed = true
					killer = false
				}

				if killer {
					killerPlayer = append(killerPlayer, v)
				}
			}

			games[game].Kill(
				strings.Join(killerPlayer, " "),
				strings.Join(killedPlayer, " "),
				payload[len(payload)-1])
		}
	}

	return games
}
