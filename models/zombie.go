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
	rand.Seed(time.Now().UnixNano())
	z.PositionX = rand.Intn(MaxX)
	z.PositionY = rand.Intn(MaxY)
	if z.PositionX == 0 {
		z.PositionX++
	}
	if z.PositionY == 0 {
		z.PositionY++
	}
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
