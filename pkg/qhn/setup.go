package qhn

import (
	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
	"github.com/kyberbits/forge/forge"
)

func Setup(runtime *forge.Runtime) (*App, error) {
	app := &App{
		runtime:       runtime,
		logger:        forge.NewLogger("app", runtime.Stdout, nil),
		hackerNewsAPI: hackernews.Client{},
	}

	config, err := buildConfig(runtime.Environment)
	if err != nil {
		return nil, err
	}

	app.config = config

	return app, nil
}
