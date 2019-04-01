package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type PongEventPayload struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}

func checkCorsOrigin(r *http.Request) bool {
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = checkCorsOrigin

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %v\n", err)
		return
	}

	defer socket.Close()

	for {
		messageType, buffer, err := socket.ReadMessage()
		if err != nil {
			break
		}

		response := Event{
			Type: "pong",
			Payload: PongEventPayload{
				Type:    messageType,
				Message: string(buffer),
			},
		}

		socket.WriteJSON(response)
	}
}
