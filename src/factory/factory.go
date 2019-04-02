package factory

import (
	"github.com/gin-gonic/gin"
	"github.com/kubesmith/kubesmith-server/src/ws"
)

type ServerFactory struct {
	hub *ws.WebsocketHub
}

type WrappedHandler func(*ServerFactory, *gin.Context)

func NewServerFactory() *ServerFactory {
	return &ServerFactory{
		hub: ws.NewWebsocketHub(),
	}
}

func (f *ServerFactory) GetHub() *ws.WebsocketHub {
	return f.hub
}

func (f *ServerFactory) WrapHandler(handler WrappedHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(f, c)
	}
}
