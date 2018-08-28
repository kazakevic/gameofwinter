package models

//MaxX - max world width
const MaxX = 10

//MaxY - max world heigth
const MaxY = 30

/*
World model
*/
type World struct {
	Players map[string]Player
	Zombies map[interface{}]Zombie
	Winner  Player
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
