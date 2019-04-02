package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/factory"
	"github.com/kubesmith/kubesmith-server/src/middleware"
	v1 "github.com/kubesmith/kubesmith-server/src/routes/v1"
)

func main() {
	// create a new router to host routes on
	router := gin.Default()

	// create a new server factory (for all of our components)
	server := factory.NewServerFactory()

	// initialize all of our middleware
	middleware.Initialize(router)

	// register the v1 routes
	v1.RegisterRoutes(router, server)

	// run the things
	go server.GetHub().Run()
	router.Run()
}
