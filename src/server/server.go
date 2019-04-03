package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/config"
	"github.com/kubesmith/kubesmith-server/src/middleware"
	"github.com/kubesmith/kubesmith-server/src/ws"
)

type Server struct {
	hub    *ws.WebsocketHub
	router *gin.Engine
}

type WrappedHandler func(*Server, *gin.Context)

func NewServer() *Server {
	// create a new server
	server := Server{
		hub:    ws.NewWebsocketHub(),
		router: gin.Default(),
	}

	// initialize all of our middleware
	middleware.Initialize(server.GetRouter())

	// return a pointer to our server struct
	return &server
}

func (s *Server) Run() {
	go s.GetHub().Run()

	port := fmt.Sprintf("0.0.0.0:%d", config.Parsed.ServerPort)
	fmt.Printf("Listening on: %s ...\n", port)
	s.GetRouter().Run(port)
}

func (s *Server) GetHub() *ws.WebsocketHub {
	return s.hub
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

func (s *Server) WrapHandler(handler WrappedHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(s, c)
	}
}
