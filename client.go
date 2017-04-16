package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan []byte
	// game is the game this client is playing in.
	game *game
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.game.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
