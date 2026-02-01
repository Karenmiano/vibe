package hub

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type liveRoom struct {
	id uuid.UUID
	clients map[*client]bool
	join chan *client
	leave chan *client
	broadcast chan *webSocketMessage
	private bool
}

func newLiveRoom(id uuid.UUID, private bool) *liveRoom {
	return &liveRoom{
		id: id,
		clients: make(map[*client]bool),
		join: make(chan *client),
		leave: make(chan *client),
		broadcast: make(chan *webSocketMessage),
		private: private,
	}
}

func (r *liveRoom) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
		case msg := <-r.broadcast:
			jsonMessage, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Unable to marshal JSON due to %s", err)
			} else {
				for client := range r.clients {
					client.receive <- jsonMessage
				}
			}
		}
	}
}