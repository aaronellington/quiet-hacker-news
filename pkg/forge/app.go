package forge

import (
	"net/http"
)

// App is the requirements for an application
type App interface {
	Handler() http.Handler
	Logger() Logger
	ListenAddress() string
}

// Run an App
func Run(app App) {
	// Create the http logger handler
	handler := &HTTPLogger{
		Log:     app.Logger(),
		Handler: app.Handler(),
	}

	// Create the server
	server := &http.Server{
		Handler: handler,
		Addr:    app.ListenAddress(),
	}

	// Log the start of the server
	app.Logger().Info(nil, "Listening and serving", map[string]interface{}{
		"address": "http://" + server.Addr,
	})

	// Listen and Serve
	if err := server.ListenAndServe(); err != nil {
		app.Logger().Error(nil, "Error During ListenAndServe", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
