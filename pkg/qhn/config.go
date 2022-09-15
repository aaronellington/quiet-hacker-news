package qhn

import "github.com/aaronellington/quiet-hacker-news/forge"

// Config is foobar
type Config struct {
	Host string
	Port int
}

func buildConfig(environment forge.Environment) (*Config, error) {
	// Defaults
	config := &Config{
		Host: "0.0.0.0",
		Port: 2222,
	}

	// Read in config variables from the environment
	if err := forge.EnvironmentUnmarshal(environment, config); err != nil {
		return nil, err
	}

	return config, nil
}
