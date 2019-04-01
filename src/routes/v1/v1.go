package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/builds"
	"github.com/kubesmith/kubesmith-server/src/routes/v1/repos"
)

func RegisterRoutes(group *gin.Engine) {
	v1 := group.Group("/v1")

	repos.RegisterRoutes(v1)
	builds.RegisterRoutes(v1)
}
