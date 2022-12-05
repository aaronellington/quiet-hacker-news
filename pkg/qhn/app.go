package qhn

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
	"github.com/aaronellington/quiet-hacker-news/resources"
	"github.com/kyberbits/forge/forge"
)

type App struct {
	runtime         *forge.Runtime
	logger          *forge.Logger
	config          *Config
	hackernewsItems []hackernews.Item
	hackerNewsAPI   hackernews.Client
}

func (app *App) ListenAddress() string {
	return fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)
}

func (app *App) Logger() *forge.Logger {
	return app.logger
}

func (app *App) Background(ctx context.Context) {
	app.runtime.KeepRunning(
		ctx,
		app,
		func(ctx context.Context) {
			app.updateCacheTick(ctx)
		},
		time.Minute*time.Duration(app.config.RefreshIntervalMinutes),
	)
}

func (app *App) Handler() http.Handler {
	return &forge.HTTPLogger{
		Logger: app.Logger().Copy("http"),
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
