package password

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/password/forgot"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/password/reset"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	password := group.Group("/password")
	forgot.RegisterRoutes(password, server)
	reset.RegisterRoutes(password, server)
}
