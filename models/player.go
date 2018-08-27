package models

type Player struct {
	Id        string
	Username  string
	PositionX int
	PositionY int
}

func (p *Player) setPosition(x, y int) {
	p.PositionX = x
	p.PositionY = y
}
