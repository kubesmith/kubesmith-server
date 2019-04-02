package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
)

func RegisterRoutes(group *gin.RouterGroup, server *factory.ServerFactory) {
	group.GET("/ws", server.WrapHandler(WebsocketUpgradeHandler))
}
