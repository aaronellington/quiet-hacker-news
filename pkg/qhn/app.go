package qhn

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronellington/quiet-hacker-news/internal/forge"
	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
	"github.com/aaronellington/quiet-hacker-news/resources"
)

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
	for range time.NewTicker(time.Minute * time.Duration(app.config.RefreshIntervalMinutes)).C {
		app.updateCacheTick()
	}
}

// Handler is foobar
func (app *App) Handler() http.Handler {
	return &forge.HTTPLogger{
		Logger: app.Logger(),
		Handler: &forge.HTTPRouter{
			Routes: map[string]http.Handler{
				"/": app.handlerRoot(),
			},
			NotFoundHandler: &forge.HTTPStatic{
				FileSystem: http.FS(resources.Public),
			},
		},
	}
}
