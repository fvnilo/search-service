package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/nylo-andry/search-service/config"
	"github.com/nylo-andry/search-service/services"
)

func main() {
	var config config.Configurations
	err := envconfig.Process("ms", &config)
	if err != nil {
		panic(err)
	}

	services.StartMicroservice(config)
}
