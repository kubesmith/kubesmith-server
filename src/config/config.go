package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var Parsed Configuration

func ParseConfig() {
	if err := envconfig.Process("kubesmith", &Parsed); err != nil {
		log.Fatal(err)
	}
}
