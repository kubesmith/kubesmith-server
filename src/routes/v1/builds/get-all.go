package builds

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type BuildsGetAllHandler struct{}

func (h *BuildsGetAllHandler) Process() ([]fixtures.Build, error) {
	return fixtures.GetBuilds(), nil
}

func BuildsGetAll(server *server.Server, c *gin.Context) {
	handler := BuildsGetAllHandler{}

	builds, err := handler.Process()
	if err != nil {
		c.Status(500)
	} else {
		c.JSON(200, builds)
	}
}
