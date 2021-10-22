package qhn

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/fuzzingbits/forge"
	"github.com/fuzzingbits/quiet-hacker-news/pkg/hackernews"
	"github.com/fuzzingbits/quiet-hacker-news/resources"
)

// NewApp builds a new App instance using the provided config
func NewApp(config Config) *App {
	app := &App{
		// Config
		listenAddress:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		pageSize:        config.PageSize,
		refreshInterval: time.Hour * config.RefreshIntervalHours,
		// Internal
		logger:        &forge.LoggerJSON{Writer: os.Stdout},
		indexTemplate: resources.Index,
		hackerNewsAPI: hackernews.Client{},
	}

	go app.startCacheUpdateLoop()

	return app
}

// App for my website
type App struct {
	listenAddress   string
	logger          forge.Logger
	pageSize        int
	refreshInterval time.Duration
	indexTemplate   *template.Template
	hackerNewsAPI   hackernews.Client
	itemCache       []hackernews.Item
}

// Handler to be used by hammer
func (app *App) Handler() http.Handler {
	router := &forge.HTTPRouter{
		Routes: map[string]http.Handler{
			"/": http.HandlerFunc(app.indexHandler),
		},
	}

	static := &forge.HTTPStatic{
		FileSystem:      http.FS(resources.Public),
		NotFoundHandler: router,
	}

	return static
}

// Logger to be used by hammer
func (app *App) Logger() forge.Logger {
	return app.logger
}

// ListenAddress to be used by hammer
func (app *App) ListenAddress() string {
	return app.listenAddress
}
