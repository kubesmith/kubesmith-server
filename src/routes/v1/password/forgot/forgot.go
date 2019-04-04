package forgot

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	forgot := group.Group("/forgot")
	forgot.POST("/", server.WrapHandler(PasswordForgotPost))
}
