package forge

import (
	"net/http"
)

// Router is foobar
type Router struct {
	Routes          map[string]http.Handler
	NotFoundHandler http.Handler
}

// ServeHTTP is foobar
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matchedRoute, found := router.Routes[r.URL.Path]
	if !found {
		correctNotFoundHandler(router.NotFoundHandler).ServeHTTP(w, r)
		return
	}

	matchedRoute.ServeHTTP(w, r)
}
