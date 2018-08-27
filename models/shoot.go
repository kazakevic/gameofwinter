package models

type Shoot struct {
	PlayerID  string
	PositionX int
	PositionY int
}

// func (s *Shoot, w *World) Status() string {
// 	var status string = "miss"

// 	for _, zombie := range w.Zombies {

// 		if s.PositionX == zombie.PositionX && s.PositionY == zombie.PositionY {
// 			status = "kill"
// 		}

// 	}

// 	return status
// }
