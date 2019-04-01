package builds

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
)

func GetAllBuilds(c *gin.Context) {
	c.JSON(200, fixtures.GetBuilds())
}
