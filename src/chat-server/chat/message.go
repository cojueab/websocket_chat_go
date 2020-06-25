package chat

import (
	"strings"
	"encoding/json"
	"fmt"
	"bytes"
	"log"
)

type Message struct {
	Text     string
	Type     int
	UserName string
	Room     string
}

func handleMessage(message []byte, c *Client){
	m:=goToJson(message)
	switch m.Type {
	case 1:
		auth := strings.Split(m.Text, ":")
		if len(auth) != 2 {
			c.send <- goToByte(Message{Text: "Ошибка", Type: -1})
		} else {
			c.name = auth[0]
			c.hub.sendUserList(c.room.name)
		}
	case 2:
		room := strings.Split(m.Text, "")
		if len(room)!=2{
			c.room = Room{name:room[0]}
		}else{
			c.room = Room{name:room[0], password:room[1]}
		}
		c.hub.sendUserList(c.room.name)
	case 3:
	case 4:
	case 5:
	default:
		m.UserName = c.name
		m.Room = c.room.name
		message = bytes.TrimSpace(bytes.Replace(goToByte(m), newline, space, -1))
		c.hub.broadcast <- message
	}
}

func goToJson(message []byte) Message{
	var m Message
	err:= json.Unmarshal(message, &m)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return m
}

func goToByte(message Message) []byte{
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	return b
}