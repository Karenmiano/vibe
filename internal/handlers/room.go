package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// holds all current clients in this room.
	clients map[*client]bool

	// join is a channel for clients wishing to join the room
	join chan *client

	// leave is a channel for clients wishing to leave the room.
	leave chan *client

	// forward is a channel that holds incoming messages that should be forwarded to the other clients.
	forward chan []byte
}

func NewRoom() *room {
	return &room{
		clients: make(map[*client]bool),
		join: make(chan *client),
		leave: make(chan *client),
		forward: make(chan []byte),
	}
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		receive: make(chan []byte, messageBufferSize),
		room: r,
	}
	r.join <- client
	defer func() { r.leave <- client}()
	go client.write()
	client.read()
}