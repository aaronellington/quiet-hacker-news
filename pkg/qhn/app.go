package qhn

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fuzzingbits/forge-wip/pkg/web"
	"github.com/fuzzingbits/quiet-hacker-news/pkg/hackernews"
)

// NewApp sets up and reteurns a new instance of App
func NewApp() *App {
	app := &App{
		Addr:            "0.0.0.0:9090",
		PageSize:        30,
		RefreshInterval: time.Hour * 1,
		indexTemplate:   template.Must(template.ParseFiles("templates/index.go.html")),
		fileSystem:      http.Dir("static"),
		hackerNewsAPI:   hackernews.Client{},
	}

	return app
}

// App is an instance of the QHN App
type App struct {
	Addr            string
	PageSize        int
	RefreshInterval time.Duration
	indexTemplate   *template.Template
	fileSystem      http.FileSystem
	itemCache       []hackernews.Item
	hackerNewsAPI   hackernews.Client
}

// Start the QHN Webserver
func (app *App) Start() error {
	go app.startCacheUpdateLoop()

	mux := http.NewServeMux()
	mux.Handle("/", &web.Handler{
		FileSystem:  app.fileSystem,
		RootHandler: http.HandlerFunc(app.indexHandler),
	})

	server := &http.Server{
		Addr:         app.Addr,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on: http://%s\n", app.Addr)
	return server.ListenAndServe()
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the template to a new bytes buffer
	buf := bytes.NewBuffer([]byte{})
	if err := app.indexTemplate.Execute(buf, app.itemCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the buffer to a byte slice
	templateBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate the CSP Header
	csp := web.GenerateContentSecurityPolicy(templateBytes, web.CSPEntries{
		Script: []string{
			"'self'",
		},
		Image: []string{
			"'self'",
		},
	})

	w.Header().Set("Content-Security-Policy", csp)
	w.Header().Set("Content-Type", "text/html")

	if _, err := w.Write(templateBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
