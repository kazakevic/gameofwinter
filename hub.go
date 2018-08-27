// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"kazakevic/gameofwinter/models"
	"regexp"
	"strconv"
	"strings"

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

				//Don't show own messages
				if message.SenderID == client.id {
					continue
				}

				select {
				case client.send <- message.Body:
					//fmt.Println("New message from", string(message[client.id]))
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
func ParseUsername(message []byte) string {
	re := regexp.MustCompile(`(?mi)^(Start)\s(.*)`)
	match := re.FindStringSubmatch(string(message))

	if len(match) > 0 {
		match[1] = strings.ToLower(match[1])
		if match[1] == "start" {
			playerName := match[2]
			return playerName
		}
	}
	return ""
}

func MakeShoot(message []byte) models.Shoot {
	re := regexp.MustCompile(`(?mi)^(Shoot)\s(\d+)\s(\d+)`)
	match := re.FindStringSubmatch(string(message))
	shoot := models.Shoot{}

	if len(match) > 0 {
		match[1] = strings.ToLower(match[1])
		if match[1] == "shoot" {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])

			shoot := models.Shoot{}
			shoot.PositionX = x
			shoot.PositionY = y

			fmt.Printf("Shoot x: %d y: %d \n", x, y)
			return shoot
		}
	}
	return shoot
}

func ParseCommands(message []byte) {

	if string(message) == "/players" {
		fmt.Print(world.GePlayersList())
	}

}
