package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/middleware"
	v1 "github.com/kubesmith/kubesmith-server/src/routes/v1"
)

func main() {
	// create a new router to host routes on
	router := gin.Default()

	// initialize all of our middleware
	middleware.Initialize(router)

	// register the v1 routes
	v1.RegisterRoutes(router)

	// run the server
	router.Run()
}
