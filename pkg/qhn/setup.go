package qhn

import (
	"encoding/json"

	"github.com/aaronellington/quiet-hacker-news/internal/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
)

// Setup is foobar
func Setup(runtime *forge.Runtime) (forge.App, error) {
	app := &App{
		runtime: runtime,
		logger: &forge.LoggerJSON{
			Encoder: json.NewEncoder(runtime.Stderr),
		},
		hackerNewsAPI: hackernews.Client{},
	}

	config, err := buildConfig(runtime.Environment)
	if err != nil {
		return nil, err
	}

	app.config = config

	return app, nil
}
