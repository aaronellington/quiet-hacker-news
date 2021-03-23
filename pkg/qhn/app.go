package qhn

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fuzzingbits/quiet-hacker-news/pkg/hackernews"
	"github.com/fuzzingbits/quiet-hacker-news/resources"
)

// NewApp sets up and reteurns a new instance of App
func NewApp() *App {
	app := &App{
		PageSize:        30,
		RefreshInterval: time.Hour * 1,
		indexTemplate:   resources.Index,
		hackerNewsAPI:   hackernews.Client{},
	}

	go app.startCacheUpdateLoop()

	return app
}

// App is an instance of the QHN App
type App struct {
	PageSize        int
	RefreshInterval time.Duration
	indexTemplate   *template.Template
	itemCache       []hackernews.Item
	hackerNewsAPI   hackernews.Client
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyBuffer := bytes.NewBuffer([]byte{})
	if err := app.indexTemplate.Execute(bodyBuffer, app.itemCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := ioutil.ReadAll(bodyBuffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
