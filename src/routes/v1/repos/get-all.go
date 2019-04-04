package repos

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/fixtures"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type ReposGetAllHandler struct{}

func (h *ReposGetAllHandler) Process() ([]fixtures.Repository, error) {
	return fixtures.GetRepos(), nil
}

func ReposGetAll(server *server.Server, c *gin.Context) {
	handler := ReposGetAllHandler{}

	repos, err := handler.Process()
	if err != nil {
		c.Status(500)
	} else {
		c.JSON(200, repos)
	}
}
