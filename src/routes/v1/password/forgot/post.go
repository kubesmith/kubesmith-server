package forgot

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/database/models"
	"github.com/kubesmith/kubesmith-server/src/server"
)

type PasswordForgotPostHandler struct{}

func (h *PasswordForgotPostHandler) Process() (*models.User, error) {
	return nil, nil
}

func PasswordForgotPost(server *server.Server, c *gin.Context) {
	c.JSON(200, true)
}
