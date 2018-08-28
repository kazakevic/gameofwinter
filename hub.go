// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"kazakevic/gameofwinter/models"
	"math/rand"
	"time"

	"github.com/rs/xid"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Inbound messages from the clients.
	//broadcast chan map[string][]byte
	broadcast chan Message
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type Message struct {
	SenderID string
	Body     []byte
	Type     string
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			client.id = xid.New().String()
			fmt.Println("New client joined ", client.id)
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Println("Client closed connection", client.id)
				//cleanup map
				world.DeletePlayer(client.id)
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				//by default show message to all
				switch message.Type {
				case "Others":
					//Don't show own messages
					if message.SenderID == client.id {
						continue
					}
				case "Self":
					//Only for me
					if message.SenderID != client.id {
						continue
					}
				}
				select {
				case client.send <- message.Body:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

/*
ParseCommands - parses other commands from messages
*/
func ParseCommands(message []byte) {
	if string(message) == "/players" {
		fmt.Print(world.GePlayersList())
	}
	if string(message) == "/online" {
		fmt.Printf("-----Online players: %d \n\n", world.GetOnlineCount())
	}
}

/*
NewMessage - broadcast message, type - Self, All, Other
*/
func (h *Hub) NewMessage(message, messageType, sender string) {
	msg := Message{SenderID: sender}
	msg.Body = []byte(message)
	msg.Type = messageType
	h.broadcast <- msg
}

/*
AnnounceZombieLoc - announces where is zombie now
*/
func (h *Hub) AnnounceZombieLoc(zombie *models.Zombie) {
	var dead = false
	for dead == false {
		time.Sleep(5000 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		zombie.ChangeLoc(rand.Intn(10), rand.Intn(30))
		x, y := zombie.GetLoc()
		s := fmt.Sprintf("WALK night-king x:%d y:%d \n", x, y)

		if zombie.Status == "dead" {
			s = "night-king is dead. This is the end for now. \n"
			dead = true
		}

		fmt.Printf(s)
		h.NewMessage(s, "All", "")
	}
}

/*
AnnounceWinner - announces who is that hero who killed zombie
*/
func (h *Hub) AnnounceWinner(world *models.World) {
	for {
		time.Sleep(10000 * time.Millisecond)

		s := "Still no winner in this battle :-( \n"

		if len(world.Winner.Username) > 1 {
			s = fmt.Sprintf("Winner now is [%s] \n", world.Winner.Username)
		}

		h.NewMessage(s, "All", "")
	}
}
