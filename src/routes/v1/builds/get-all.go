package builds

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type GetAllBuildsHandler struct{}

func (h *GetAllBuildsHandler) Process() ([]fixtures.Build, error) {
	return fixtures.GetBuilds(), nil
}

func GetAllBuilds(server *server.Server, c *gin.Context) {
	handler := GetAllBuildsHandler{}

	builds, err := handler.Process()
	if err != nil {
		c.Status(500)
	} else {
		c.JSON(200, builds)
	}
}
