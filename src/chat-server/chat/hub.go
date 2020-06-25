package chat

import (
	"encoding/json"
)

type Mans struct {
	Names []string `json:"names"`
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.sendUserList(client.room.name)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.sendUserList(client.room.name)
		case message := <-h.broadcast:
			h.broadSend(message)
		}
	}
}

func (h *Hub) list() string {
	names := make([]string, len(h.clients))
	i := 0
	for client := range h.clients {
		names[i] = client.name
		i++
	}
	res2D := &Mans{
		Names: names}
	res2B, _ := json.Marshal(res2D)
	return string(res2B)
}
