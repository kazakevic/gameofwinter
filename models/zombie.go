package models

type Zombie struct {
	PositionX int
	PositionY int
	Status    string
}

/*
GetLoc gets zombie location
*/
func (z *Zombie) GetLoc() (int, int) {
	return z.PositionX, z.PositionY
}

/*
ChangeLoc changes zombie location
*/
func (z *Zombie) ChangeLoc(x, y int) {
	z.PositionX = x
	z.PositionY = y
	if z.PositionX == 0 {
		z.PositionX++
	}
	if z.PositionY == 0 {
		z.PositionY++
	}
}

/*
Hit calculates collision between zombie and shoot
*/
func (z *Zombie) Hit(shoot Shoot) bool {

	if (z.PositionX == shoot.PositionX) && (z.PositionY == shoot.PositionY) {
		z.Status = "dead"
		return true
	}

	return false
}
