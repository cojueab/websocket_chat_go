package chat

import (
	"net/http"
	"log"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub: hub,
		conn: conn,
		send: make(chan []byte, 256),
		host: r.RemoteAddr,
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
