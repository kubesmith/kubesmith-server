package services

import (
	"fmt"

	"github.com/carldanley/go-services"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/kubesmith/kubesmith-server/src/config"
)

var Factory *services.Factory

func Setup() {
	if Factory != nil {
		return
	}

	Factory = services.NewFactory()

	registerDatabase()

	Factory.Subscribe(configureDatabase)
	Factory.Subscribe(logServiceEvents)
}

func GetDB() *gorm.DB {
	svc, _ := Factory.Get(services.ServiceTypeGorm)
	return svc.GetClient().(*gorm.DB)
}

func registerDatabase() {
	cfg := services.Config{
		Host:     config.Parsed.DatabaseHost,
		Port:     config.Parsed.DatabasePort,
		Username: config.Parsed.DatabaseUser,
		Password: config.Parsed.DatabasePass,
		Database: config.Parsed.DatabaseName,

		ReconnectEnabled: true,
	}

	Factory.Register(services.ServiceTypeGorm, cfg)
}

func configureDatabase(event services.Event) {
	if event.Code != services.ServiceHealthy || event.ServiceType != services.ServiceTypeGorm {
		return
	}

	client := GetDB()

	// set debugging for the database connection
	client.LogMode(config.Parsed.DatabaseDebug)

	// setup the connection settings
	client.DB().SetMaxIdleConns(config.Parsed.DatabaseMaxIdle)
	client.DB().SetMaxOpenConns(config.Parsed.DatabaseMaxOpen)
}

func logServiceEvents(event services.Event) {
	var status string

	cyan := color.New(color.FgHiCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()

	switch event.Code {
	case services.ServiceUnhealthy:
		status = red("unhealthy")
	case services.ServiceHealthy:
		status = green("healthy")
	case services.ServiceConnected:
		status = green("connected")
	case services.ServiceDisconnected:
		status = red("disconnected")
	case services.ServiceReconnecting:
		status = yellow("reconnecting")
	case services.ServiceReconnected:
		status = green("reconnected")
	case services.ServiceCouldNotConnect:
		status = red("could not connect")
	}

	fmt.Printf("%s: %s\n", cyan(event.ServiceType), status)
}
