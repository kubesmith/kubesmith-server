package reset

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	reset := group.Group("/reset")
	reset.POST("/", server.WrapHandler(PasswordResetPost))
}
