package hub

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	hub *Hub
	rooms map[*liveRoom]bool
	receive chan []byte
}

const messageBufferSize = 256

func newClient(socket *websocket.Conn, hub *Hub) *client {
	return &client{
		socket: socket,
		hub: hub,
		receive: make(chan []byte, messageBufferSize),
	}
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, jsonMessage, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.handleNewMessage(jsonMessage)
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil{
			return
		}
	}
}

func (c *client) disconnect() {
	c.hub.leave <- c

	for room := range c.rooms {
		room.leave <- c
	}
	close(c.receive)
}

func (c *client) handleNewMessage(jsonMessage []byte) {
	var wsMessage webSocketMessage

	err := json.Unmarshal(jsonMessage, &wsMessage)
	if err != nil {
		log.Printf("Unable to unmarshal JSON due to %s", err)
		return
	}

	wsMessage.Sender = c

	switch wsMessage.Action {
	case sendMessageAction:
		roomID := wsMessage.Target
		if room := c.hub.findRoomByID(roomID); room != nil {
			room.broadcast <- &wsMessage
		} else {
			// wake up room
			// add all clients that belong to it and are in hub
			// broadcast message
		}
	case joinRoomAction:
		// save that state in db
		// if room is online add client to that room
	}
}
