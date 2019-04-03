package repos

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	repos := group.Group("/repos")
	repos.GET("/", server.WrapHandler(GetAllRepos))
}
