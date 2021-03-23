package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fuzzingbits/forge"
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
	configuration, err := getConfiguration()
	if err != nil {
		panic(err)
	}

	// Setup the server
	server := http.Server{
		Addr:         fmt.Sprintf("%s:%s", configuration.Host, configuration.Port),
		Handler:      getHandler(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Printf("Listening on http://%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getConfiguration() (*Configuration, error) {
	// Setup defaults
	configuration := &Configuration{
		Host: "0.0.0.0",
		Port: "8000",
	}

	if err := forge.ParseEnvironment(configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func getHandler() http.Handler {
	app := qhn.NewApp()

	// Build the primary router
	router := &forge.Router{
		// NotFoundHander: app,
	}

	router.Handle("/", app)

	// Configure static file serving
	static := &forge.Static{
		FileSystem:      http.FS(resources.Public),
		NotFoundHandler: router,
	}

	// Configure the logger
	logger := &forge.Logger{
		Handler: static,
	}

	return logger
}
