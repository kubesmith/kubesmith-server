package main

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/kubesmith/kubesmith-server/src/config"
	v1 "github.com/kubesmith/kubesmith-server/src/routes/v1"
	"github.com/kubesmith/kubesmith-server/src/server"
)

var signalChannel chan os.Signal

type sqlLogger struct{}

func (s sqlLogger) Print(args ...interface{}) {}

func handleInterrupts() {
	for range signalChannel {
		os.Exit(1)
	}
}

func setupEnvironment() {
	// initialize the gin package
	gin.SetMode(gin.ReleaseMode)

	// set GOMAXPROCS
	runtime.GOMAXPROCS(1)

	// set the logger for mysql
	mysql.SetLogger(sqlLogger{})

	// setup the random seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// make the signal channel and register it
	signalChannel = make(chan os.Signal, 1)

	// register the interrupt signals to be sent to the signal channel
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)

	// handle any interrupts that get caught
	go handleInterrupts()

	// get the config
	config.ParseConfig()

	// setup the environment
	setupEnvironment()

	// create a new server
	server := server.NewServer()

	// register our routes
	v1.RegisterRoutes(server.GetRouter(), server)

	// run the server
	server.Run()
}
