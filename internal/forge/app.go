package forge

import (
	"fmt"
	"net/http"
)

// App is foobar
type App interface {
	// The Logger to be used
	Logger() Logger
}

// WebApp is foobar
type WebApp interface {
	App
	// How it responds to web requests
	Handler() http.Handler
	// What address the webserver should listen on
	ListenAddress() string
	// Tasks to run in the background (a Go routine)
	Background()
}

// Run is foobar
func Run(app WebApp) error {
	go app.Background()

	httpServer := &http.Server{
		Addr:     app.ListenAddress(),
		Handler:  app.Handler(),
		ErrorLog: app.Logger().StandardLogger(),
	}

	app.Logger().Info("Serving web application", map[string]interface{}{
		"listen": fmt.Sprintf("http://%s", app.ListenAddress()),
	})

	if err := httpServer.ListenAndServe(); err != nil {
		app.Logger().Critical("Webserver failed to start", map[string]interface{}{
			"addr": app.ListenAddress(),
			"err":  err,
		})
		return err
	}

	return nil
}
