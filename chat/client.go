package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err != nil {
			break
		} else {
			c.room.forward <- msg
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
