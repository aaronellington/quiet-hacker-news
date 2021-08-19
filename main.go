package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fuzzingbits/forge"
	"github.com/fuzzingbits/forge/hammer"
	"github.com/fuzzingbits/forge/workbench"
	"github.com/fuzzingbits/quiet-hacker-news/pkg/qhn"
	"github.com/fuzzingbits/quiet-hacker-news/resources"
)

// Configuration is the structure of the configuration options
type Configuration struct {
	Host string
	Port string `env:"PORT"`
}

func main() {
	// Get the configuration
	configuration := getConfiguration()

	// Build the hammer.App
	app := &hammer.App{
		ListenAddress: fmt.Sprintf("%s:%s", configuration.Host, configuration.Port),
		Logger:        &workbench.LoggerJSON{Writer: os.Stdout},
		Handler: &forge.Router{
			Routes: map[string]http.Handler{
				"/": qhn.NewApp(),
			},
		},
		Middleware: []forge.Middleware{
			&forge.Static{
				FileSystem: http.FS(resources.Public),
			},
		},
	}

	app.Run()
}

func getConfiguration() *Configuration {
	// Setup defaults
	configuration := &Configuration{
		Host: "0.0.0.0",
		Port: "8000",
	}

	if err := forge.ParseEnvironment(configuration); err != nil {
		panic(err)
	}

	return configuration
}
