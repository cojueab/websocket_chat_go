package routes

import (
	"chat-server/chat"
	"net/http"
)

func activeSocket(hub *chat.Hub){
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
}

