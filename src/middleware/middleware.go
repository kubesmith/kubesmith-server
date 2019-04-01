package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	router.Use(cors.Default())
}
