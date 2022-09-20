package main

import (
	"github.com/aaronellington/quiet-hacker-news/internal/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/qhn"
)

func main() {
	runtime := forge.NewRuntime()

	if err := runtime.ReadInDefaultEnvironmentFiles(); err != nil {
		panic(err)
	}

	app, err := qhn.Setup(runtime)
	if err != nil {
		panic(err)
	}

	forge.Run(app)
}
