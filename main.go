package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fuzzingbits/quiet-hacker-news/pkg/qhn"
)

func main() {
	app := qhn.NewApp()

	// If a custom port number is specified, override the default
	if port := os.Getenv("PORT"); port != "" {
		app.Addr = fmt.Sprintf("0.0.0.0:%s", port)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
