package forge

import (
	"net/http"
)

// HTTPRouter is foobar
type HTTPRouter struct {
	Routes          map[string]http.Handler
	NotFoundHandler http.Handler
}

// ServeHTTP is foobar
func (httpRouter *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matchedRoute, found := httpRouter.Routes[r.URL.Path]
	if !found {
		correctNotFoundHandler(httpRouter.NotFoundHandler).ServeHTTP(w, r)
		return
	}

	matchedRoute.ServeHTTP(w, r)
}
