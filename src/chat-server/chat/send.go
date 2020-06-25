package chat


func (h *Hub) broadSend(b []byte) {
	m:=goToJson(b)
	for client := range h.clients {
		if client.room.name == m.Room {
			select {
			case client.send <- b:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

func (h *Hub) sendUserList(room string) {
	h.broadSend(goToByte(Message{Text: h.list(), Type: 2, Room: room}))
}
