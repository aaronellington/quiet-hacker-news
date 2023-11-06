package main

import (
	"context"

	"github.com/aaronellington/quiet-hacker-news/pkg/qhn"
	"github.com/kyberbits/forge/forge"
)

func main() {
	runtime := forge.NewRuntime()

	app, err := qhn.Setup(runtime)
	if err != nil {
		panic(err)
	}

	_ = runtime.Serve(context.Background(), app)
}
