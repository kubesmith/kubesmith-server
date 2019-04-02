package repos

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
)

func RegisterRoutes(group *gin.RouterGroup, server *factory.ServerFactory) {
	repos := group.Group("/repos")
	repos.GET("/", server.WrapHandler(GetAllRepos))
}
