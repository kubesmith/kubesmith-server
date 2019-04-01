package builds

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	builds := group.Group("/builds")
	builds.GET("/", GetAllBuilds)
}
