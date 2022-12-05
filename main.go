package main

import (
	"context"

	"github.com/aaronellington/quiet-hacker-news/pkg/qhn"
	"github.com/kyberbits/forge/forge"
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

	runtime.Serve(context.Background(), app)
}
