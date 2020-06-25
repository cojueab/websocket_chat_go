package routes

import "chat-server/chat"

func activeChat(hub *chat.Hub){
	hub.Run()
}