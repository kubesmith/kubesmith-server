package ws

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// Maximum message size allowed from peer.
	MaxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebsocketHubClient struct {
	identity string
	socket   *websocket.Conn
	messages chan []byte
	hub      *WebsocketHub
}

type WebsocketPayload struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func NewWebsocketHubClient(socket *websocket.Conn, hub *WebsocketHub) *WebsocketHubClient {
	return &WebsocketHubClient{
		identity: "",
		socket:   socket,
		messages: make(chan []byte),
		hub:      hub,
	}
}

func (c *WebsocketHubClient) Send(message []byte) {
	c.messages <- message
}

func (c *WebsocketHubClient) GetIdentity() string {
	return c.identity
}

func (c *WebsocketHubClient) SetIdentity(identity string) {
	c.identity = identity
}

func (c *WebsocketHubClient) Run() {
	go c.startPings()
	go c.readMessages()
	go c.writeMessages()
}

func (c *WebsocketHubClient) startPings() {
	ticker := time.NewTicker(PingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.socket.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *WebsocketHubClient) readMessages() {
	defer func() {
		c.hub.Unregister(c)
		c.socket.Close()
	}()

	c.socket.SetReadLimit(MaxMessageSize)
	c.socket.SetReadDeadline(time.Now().Add(PongWait))
	c.socket.SetPongHandler(func(string) error {
		c.socket.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		c.parseMessage(bytes.TrimSpace(bytes.Replace(message, newline, space, -1)))
	}
}

func (c *WebsocketHubClient) writeMessages() {
	defer c.socket.Close()

	for {
		select {
		case message, ok := <-c.messages:
			c.socket.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := c.socket.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			writer.Write(message)
			count := len(c.messages)
			for i := 0; i < count; i++ {
				writer.Write(newline)
				writer.Write(<-c.messages)
			}

			if err := writer.Close(); err != nil {
				return
			}
		}
	}
}

func (c *WebsocketHubClient) parseMessage(buffer []byte) {
	payload := WebsocketPayload{}
	if err := json.Unmarshal(buffer, &payload); err != nil {
		return
	}

	// spew.Dump(payload)
}
