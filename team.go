package main

import (
	"fmt"
	"log"
	"math/rand"
)

const TeamPlayerCount int = 5

type Team struct {
	Name         string
	Point        int
	OnAttack     bool
	AttackCount  int
	BallHolder   *Player
	PassorPlayer *Player
	log          *log.Logger

	Players []*Player
}

func NewTeam(name string, rnd *rand.Rand, log *log.Logger) *Team {
	t := &Team{
		Name: name,
		log:  log,
	}
	t.Players = make([]*Player, TeamPlayerCount)

	for i := 0; i < TeamPlayerCount; i++ {
		t.Players[i] = NewPlayer(fmt.Sprintf("%s-%d", name, i), rnd.Intn(10), i, log)
	}
	return t
}

func (t *Team) LooseBall() {
	t.OnAttack = false
}

func (t *Team) Scores(point int, player *Player, assist *Player) {
	t.log.Printf("[Team] - fn:Scores - Score calculating ")
	t.Point = t.Point + point
	player.Score = player.Score + point
	assist.Assist += 1
}

func (t *Team) GetTopScorer() *Player {
	player := t.Players[0]
	for _, p := range t.Players {
		if p.Score >= player.Score {
			player = p
		}
	}
	t.log.Printf("[Team] - fn:GetTopScorer - Scor:%v", player.Score)

	return player
}

func (t *Team) GetTopAssist() *Player {

	player := t.Players[0]
	for _, p := range t.Players {
		if p.Assist >= player.Assist {
			player = p
		}
	}
	t.log.Printf("[Team] - fn:GetTopAssist - Assist:%v", player.Assist)
	return player
}

func (t *Team) SimulateAttack(ballHolder *Player, rnd *rand.Rand) {
	t.log.Printf("[Team] - fn:SimulateAttack")

	t.BallHolder = ballHolder
	for i := 0; i < 5; i++ {
		lastPassor := t.BallHolder.UniformNumber
		newHolderNumber := t.BallHolder.Pass(rnd.Intn(5), rnd)

		t.BallHolder = t.Players[newHolderNumber]
		t.PassorPlayer = t.Players[lastPassor]
	}
	point := t.BallHolder.Shoot(rnd.Intn(100))
	t.log.Printf("[Team] - fn:SimulateAttack - point:%d", point)

	if point != 0 {
		t.Scores(point, t.BallHolder, t.PassorPlayer)
	} else {
		t.OnAttack = false
	}
	t.AttackCount++
	t.log.Printf("[Team] - fn:SimulateAttack - AttackCount:%d", t.AttackCount)

}
