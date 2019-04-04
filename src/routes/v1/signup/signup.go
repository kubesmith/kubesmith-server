package signup

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/server"
)

func RegisterRoutes(group *gin.RouterGroup, server *server.Server) {
	signup := group.Group("/signup")
	signup.POST("/", server.WrapHandler(SignupPost))
}
