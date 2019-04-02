package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kubesmith/kubesmith-server/src/factory"
	"github.com/kubesmith/kubesmith-server/src/ws"
)

func checkCorsOrigin(r *http.Request) bool {
	return true
}

func WebsocketUpgradeHandler(server *factory.ServerFactory, c *gin.Context) {
	// create a new websocket upgrader
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// setup cors for the websocket upgrader
	upgrader.CheckOrigin = checkCorsOrigin

	// attempt to upgrade the connection to a websocket using the upgrader
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %v\n", err)
		return
	}

	// get the hub
	hub := server.GetHub()

	// create a new websocket hub client
	client := ws.NewWebsocketHubClient(socket, hub)

	// fake an identity
	client.SetIdentity("kubesmith")

	// register this client with the hub
	hub.Register(client)

	// run the client
	go client.Run()
}
