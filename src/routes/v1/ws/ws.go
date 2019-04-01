package ws

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/ws", func(c *gin.Context) {
		handler(c.Writer, c.Request)
	})
}
