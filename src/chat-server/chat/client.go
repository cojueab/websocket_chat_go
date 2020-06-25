package chat

import (
	"github.com/gorilla/websocket"

	"time"
	"log"
)

type Room struct {
	name string
	password string
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	host string
	name string
	room Room
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait));
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		handleMessage(message, c)
	}
}