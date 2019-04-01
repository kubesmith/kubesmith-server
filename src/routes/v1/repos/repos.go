package repos

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	repos := group.Group("/repos")
	repos.GET("/", GetAllRepos)
}
