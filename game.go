package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	DurationSec  = 240
	AttackMaxSec = 24
)

type Game struct {
	Name string

	Teams           []*Team
	TopScorer       *Player
	TopAssistPlayer *Player
	AttackerTeam    int

	rand *rand.Rand
	log  *log.Logger
	sync.Mutex
	cancelFunc context.CancelFunc
}

func NewGame(name string) *Game {
	game := Game{
		Name: name,
		log:  log.New(os.Stdout, name, log.Ltime),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	game2 := Game{
		Name: name,
		log:  log.New(os.Stdout, name, log.Ltime),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	game.Teams = make([]*Team, 2)
	game.Teams[0] = NewTeam(name+":Black", game.rand, game.log)
	game.Teams[1] = NewTeam(name+":White", game2.rand, game2.log)

	return &game
}

func (g *Game) Start() {
	g.log.Printf("Starting %s\n", g.Name)
	var ctx context.Context
	ctx, g.cancelFunc = context.WithTimeout(context.Background(), time.Second*DurationSec)

	timer := time.NewTicker(time.Second * time.Duration(g.rand.Intn(AttackMaxSec)+1))

	g.AttackerTeam = g.rand.Intn(2)
	for {
		select {
		case <-timer.C:
			g.log.Printf("new possition")
			g.DetermineAttack()
		case <-ctx.Done():
			timer.Stop()
			return
		}
	}
}

func (g *Game) DetermineAttack() {
	g.Lock()
	defer g.Unlock()
	team := g.Teams[g.AttackerTeam]
	// g.log.Printf("[Game] - Team:%s", team.Name)

	team.SimulateAttack(team.Players[g.rand.Intn(5)], g.rand)
	g.AttackerTeam = (g.AttackerTeam + 1) % 2
	fmt.Println("AtackerTeam", g.AttackerTeam)
	g.updateData()
}

func (g *Game) End() {
	g.log.Printf("Canceling %s\n", g.Name)
	g.cancelFunc()
}

func (g *Game) updateData() {
	p := g.getScorerPlayer()
	if p != nil {
		if g.TopScorer == nil {
			g.TopScorer = p
		} else if p.Score > g.TopScorer.Score {
			g.TopScorer = p
		}
	}

	p = g.getMostAssistPlayer()
	if p != nil {
		if g.TopAssistPlayer == nil {
			g.TopAssistPlayer = p
		} else if p.Assist > g.TopAssistPlayer.Assist {
			g.TopAssistPlayer = p
		}
	}

	fmt.Println("TopAssist:", g.TopAssistPlayer.Name, g.TopAssistPlayer.Assist)
	fmt.Println("TopScore:", g.TopScorer.Name, g.TopScorer.Score)

}

func (g *Game) getMostAssistPlayer() (player *Player) {
	for _, t := range g.Teams {
		p := t.GetTopAssist()
		if player == nil || p.Assist > player.Assist {
			player = p
		}
	}
	return
}
func (g *Game) getScorerPlayer() (player *Player) {
	for _, t := range g.Teams {
		p := t.GetTopScorer()
		if player == nil || p.Score > player.Score {
			player = p
		}
	}
	return
}
