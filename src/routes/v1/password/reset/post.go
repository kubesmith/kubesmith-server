package reset

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/database/models"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type PasswordResetPostHandler struct{}

func (h *PasswordResetPostHandler) Process() (*models.User, error) {
	return nil, nil
}

func PasswordResetPost(server *server.Server, c *gin.Context) {
	c.JSON(200, true)
}
