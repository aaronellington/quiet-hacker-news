package qhn

import (
	"net/http"

	"github.com/aaronellington/quiet-hacker-news/resources"
	"github.com/kyberbits/forge/forge"
)

type APIResponse struct {
	Message string `json:"message"`
}

func (app *App) handlerRoot() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerContext := forge.NewHandlerContext(w, r)
		handlerContext.ExecuteTemplate(resources.Index, app.hackernewsItems)
	})
}
