package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func init() {
	routes = append(routes, Route{"serveWs", "GET", "/ws", serveWs})
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}
