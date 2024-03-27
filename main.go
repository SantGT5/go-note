package main

import (
	"github.com/SantGT5/gopportunities/config"
	"github.com/SantGT5/gopportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	// init Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config init error: %v", err.Error())
		return
	}

	// init Routers
	router.Initialize()
}
