package routes

import (
	"chat-server/chat"
)

func ActiveRoutes(hub *chat.Hub)  {
	activeChat(hub)
	activeStatic()
	activeSocket(hub)
}

