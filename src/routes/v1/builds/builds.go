package builds

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	builds := group.Group("/builds")
	builds.GET("/", server.WrapHandler(GetAllBuilds))
}
