package qhn

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronellington/quiet-hacker-news/internal/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
	"github.com/aaronellington/quiet-hacker-news/resources"
)

const refreshInterval = time.Hour * 1
const pageSize = 30

// App is foobar
type App struct {
	runtime         *forge.Runtime
	logger          forge.Logger
	config          *Config
	hackernewsItems []hackernews.Item
	hackerNewsAPI   hackernews.Client
}

// ListenAddress is foobar
func (app *App) ListenAddress() string {
	return fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)
}

// Logger is foobar
func (app *App) Logger() forge.Logger {
	return app.logger
}

// Background is foobar
func (app *App) Background() {
	app.updateCacheTick()
	for range time.NewTicker(refreshInterval).C {
		app.updateCacheTick()
	}
}

// Handler is foobar
func (app *App) Handler() http.Handler {
	return &forge.Router{
		Routes: map[string]http.Handler{
			"/": app.handlerRoot(),
		},
		NotFoundHandler: &forge.Static{
			FileSystem: http.FS(resources.Public),
		},
	}
}
