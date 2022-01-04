package main

import (
	"github.com/aaronellington/quiet-hacker-news/pkg/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/qhn"
)

func main() {
	config := getConfig()
	app := qhn.NewApp(config)
	forge.Run(app)
}

func getConfig() qhn.Config {
	// Set defaults
	config := qhn.Config{
		Host:                 "0.0.0.0",
		Port:                 8000,
		PageSize:             30,
		RefreshIntervalHours: 1,
	}

	if err := forge.Env(&config); err != nil {
		panic(err)
	}

	return config
}
