package qhn

import (
	"net/http"
	"time"

	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
	"github.com/kyberbits/forge/forge"
)

type TemplatePayload struct {
	Items  []hackernews.Item
	Uptime string
}

func (app *App) handlerRoot() http.Handler {
	now := time.Now()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerContext := forge.NewHandlerContext(w, r)
		handlerContext.ExecuteTemplate(app.resources.Index, TemplatePayload{
			Items:  app.hackernewsItems,
			Uptime: time.Since(now).Round(time.Second).String(),
		})
	})
}
