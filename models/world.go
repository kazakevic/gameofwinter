package models

import (
	"math/rand"
	"time"
)

/*
World model
*/
type World struct {
	Players map[string]Player
	Zombies map[interface{}]Zombie
	Winner  Player
	MaxX    int
	MaxY    int
}

func NewWorld(x, y int) *World {
	w := new(World)
	w.MaxX = x
	w.MaxY = y
	w.Players = make(map[string]Player)
	return w
}

func (w *World) AddPlayers(player Player) {
	w.Players[player.Id] = player
}

/*
SpawnPlayer - add new player to the world
*/
func (w *World) SpawnPlayer(id, username string) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(w.MaxX)
	player := Player{Id: id, Username: username, PositionX: x, PositionY: w.MaxY}
	w.AddPlayers(player)
}

func (w *World) DeletePlayer(playerId string) {
	delete(w.Players, playerId)
}

/*
GetOnlineCount  Joined players count
*/
func (w *World) GetOnlineCount() int {
	online := len(w.Players)
	return online
}

func (w *World) GePlayersList() string {
	str := ""
	for id, player := range w.Players {
		str += "Player ID: [" + id + "] Player name: [" + player.Username + "] \n"
	}
	return str
}
