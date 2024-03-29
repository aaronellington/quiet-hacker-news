package qhn

import (
	"github.com/aaronellington/environment-go/environment"
)

type Config struct {
	Host                   string `env:"HOST"`
	Port                   int    `env:"PORT"`
	RefreshIntervalMinutes int    `env:"QHN_REFRESH_INTERVAL_MINUTES"`
	PageSize               int    `env:"QHN_PAGE_SIZE"`
}

func buildConfig(environment *environment.Environment) (*Config, error) {
	// Defaults
	config := &Config{
		Host:                   "0.0.0.0",
		Port:                   2222,
		RefreshIntervalMinutes: 60,
		PageSize:               30,
	}

	// Read in config variables from the environment
	if err := environment.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
