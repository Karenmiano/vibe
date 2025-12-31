package hub

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*client]bool

	roomsMu sync.RWMutex
	rooms map[uuid.UUID]*liveRoom

	join chan *client
	leave chan *client
	broadcast chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*client]bool),
		rooms: make(map[uuid.UUID]*liveRoom),
		join: make(chan *client),
		leave: make(chan *client),
		broadcast: make(chan []byte),
	}
}


func (h *Hub) Run() {
	for {
		select {
		case client := <-h.join:
			h.clients[client] = true
		case client := <-h.leave:
			delete(h.clients, client)
		case msg := <-h.broadcast:
			for client := range h.clients {
				client.receive <- msg
			}
		}
	}
}

const socketBufferSize = 1024

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	client := newClient(socket, h)
	h.join <- client
	go client.write()
	client.read()

	defer client.disconnect()
}

func (h *Hub) findRoomByID(ID uuid.UUID) *liveRoom {
	h.roomsMu.RLock()
	defer h.roomsMu.RUnlock()
	return h.rooms[ID]
}
