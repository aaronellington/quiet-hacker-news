package qhn

import (
	"bytes"
	"io"
	"net/http"

	"github.com/aaronellington/quiet-hacker-news/forge"
	"github.com/aaronellington/quiet-hacker-news/resources"
)

// APIResponse is foobar
type APIResponse struct {
	Message string `json:"message"`
}

func (app *App) handlerRoot() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBuffer := bytes.NewBuffer([]byte{})
		if err := resources.Index.Execute(bodyBuffer, app.hackernewsItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBytes, err := io.ReadAll(bodyBuffer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		forge.RespondHTML(w, http.StatusOK, string(responseBytes))
	})
}
