package models

import (
	"math/rand"
	"time"
)

type Player struct {
	Id        string
	Username  string
	PositionX int
	PositionY int
}

func (p *Player) Spawn() {

	rand.Seed(time.Now().UnixNano())
	p.PositionX = rand.Intn(10)
	p.PositionY = 30
	if p.PositionX == 0 {
		p.PositionX++
	}
	if p.PositionY == 0 {
		p.PositionY++
	}

}

func (p *Player) setPosition(x, y int) {
	p.PositionX = x
	p.PositionY = y
}
