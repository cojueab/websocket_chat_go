package main

import (
	"flag"
	"log"
	"net/http"
	"chat-server/chat"
	"chat-server/routes"
)

var port = flag.String("port", ":8080", "port for chat")

func main() {
	flag.Parse()
	routes.ActiveRoutes(chat.NewHub())
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}