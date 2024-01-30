package player

import "sort"

func RankPlayers(players map[string]int) PlayerList {
	pl := make(PlayerList, len(players))
	i := 0
	for k, v := range players {
		pl[i] = Player{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Player struct {
	Player string
	Kills  int
}

type PlayerList []Player

func (p PlayerList) Len() int           { return len(p) }
func (p PlayerList) Less(i, j int) bool { return p[i].Kills < p[j].Kills }
func (p PlayerList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
