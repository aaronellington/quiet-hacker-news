package main

import (
	"github.com/aaronellington/quiet-hacker-news/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/qhn"
)

func main() {
	runtime := forge.NewRuntime()

	if err := forge.EnvironmentReadFromDefaultFiles(runtime.Environment); err != nil {
		panic(err)
	}

	app, err := qhn.Setup(runtime)
	if err != nil {
		panic(err)
	}

	forge.Run(app)
}
