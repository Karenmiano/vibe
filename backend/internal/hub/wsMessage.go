package hub

import (
	"github.com/google/uuid"
)

const sendMessageAction = "send-message"
const joinRoomAction = "join-room"

type webSocketMessage struct {
	Action string `json:"action"`
	Message string `json:"message"`
	Target uuid.UUID `json:"target"`
	Sender *client `json:"sender"`
}