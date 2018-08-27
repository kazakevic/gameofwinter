package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Zombie struct {
	PositionX int
	PositionY int
}

func (z *Zombie) GetLoc() (x, y int) {
	return z.PositionX, z.PositionY
}

func (z *Zombie) ChangeLoc() {
	z.PositionX = rand.Intn(10)
	z.PositionY = rand.Intn(30)
	fmt.Printf("Spawned night-king at x: %d y: %d \n", z.PositionX, z.PositionY)

	for {
		time.Sleep(5000 * time.Millisecond)
		z.PositionX = rand.Intn(10)
		z.PositionY = rand.Intn(30)
		fmt.Printf("night-king is now x: %d y: %d \n", z.PositionX, z.PositionY)
	}
}

func (z *Zombie) Hit(shoot Shoot) bool {

	if (z.PositionX == shoot.PositionX) && (z.PositionY == shoot.PositionY) {
		return true
	}

	return false
}
