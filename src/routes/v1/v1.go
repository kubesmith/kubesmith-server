package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/builds"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/repos"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/ws"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.Engine, server *server.Server) {
	v1 := group.Group("/v1")

	ws.RegisterRoutes(v1, server)
	repos.RegisterRoutes(v1, server)
	builds.RegisterRoutes(v1, server)
}
