package models

//MaxX - max world width
const MaxX = 10

//MaxY - max world heigth
const MaxY = 30

/*
World model
*/
type World struct {
	Players map[interface{}]Player
	Zombies map[interface{}]Zombie
}

/*
GetOnlineCount  Joined players count
*/
func (w *World) GetOnlineCount() int {
	online := len(w.Players)
	return online
}
