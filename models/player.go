package models

type Player struct {
	Username  string
	PositionX int
	PositionY int
}

func (p *Player) setPosition(x, y int) {
	p.PositionX = x
	p.PositionY = y
}
