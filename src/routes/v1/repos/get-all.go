package repos

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
)

func GetAllRepos(server *factory.ServerFactory, c *gin.Context) {
	c.JSON(200, fixtures.GetRepos())
}
