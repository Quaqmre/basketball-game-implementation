package main

import (
	"fmt"
	"log"
	"math/rand"
)

type Player struct {
	Name          string
	Assist        int
	Score         int
	Ability       int
	BallHolder    bool
	UniformNumber int

	log *log.Logger
}

func NewPlayer(name string, ability int, number int, log *log.Logger) *Player {
	return &Player{
		Name:          name,
		Ability:       ability,
		UniformNumber: number,
		log:           log,
	}
}

func (p *Player) Shoot(chance int) int {
	fmt.Println("shooting")
	switch (p.Ability * chance) % 2 {
	case 1:
		return 3
	default:
		return 2
	}
}

//Pass should take between 0-5 numbers
func (p *Player) Pass(player int, rnd *rand.Rand) int {
	if p.UniformNumber != player {
		p.BallHolder = false
		return player
	}

	return p.Pass(rnd.Intn(5), rnd)

}

func (p *Player) String() string {
	return p.Name
}
