package models

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

func (w *World) GetSize() (int, int) {
	return w.MaxX, w.MaxY
}

func (w *World) AddPlayers(player Player) {
	w.Players[player.Id] = player
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
