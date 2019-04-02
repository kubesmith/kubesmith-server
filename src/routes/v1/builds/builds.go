package builds

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
)

func RegisterRoutes(group *gin.RouterGroup, server *factory.ServerFactory) {
	builds := group.Group("/builds")
	builds.GET("/", server.WrapHandler(GetAllBuilds))
}
