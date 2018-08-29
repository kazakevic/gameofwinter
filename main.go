// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"kazakevic/gameofwinter/models"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

//Globals
var zombie = models.Zombie{Status: "alive"}

//Create world 10x30
var world = models.NewWorld(10, 30)

func main() {

	flag.Parse()
	hub := newHub()
	go hub.run()
	go hub.AnnounceZombieLoc(&zombie)
	go hub.AnnounceWinner(world)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
