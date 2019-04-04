package signup

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	validate = validator.New()
	signup := group.Group("/signup")
	signup.POST("/", server.WrapHandler(SignupPost))
}
