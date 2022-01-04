package forge

import (
	"net/http"
	"strings"
)

// HTTPRouter serves http.Requests for a predefined map of Paths
type HTTPRouter struct {
	NotFoundHandler http.Handler
	Routes          map[string]http.Handler
}

// ServerHTTP satisfies the http.Handler interface
func (router *HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if stripTrailingSlash(w, r) {
		return
	}

	matchingRoute, found := router.Routes[r.URL.Path]
	if !found {
		if router.NotFoundHandler != nil {
			router.NotFoundHandler.ServeHTTP(w, r)
			return
		}

		notFoundHandler(w, r)
		return
	}

	if matchingRoute != nil {
		matchingRoute.ServeHTTP(w, r)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	RespondText(w, http.StatusNotFound, []byte(ResponseTextNotFound))
}

func stripTrailingSlash(w http.ResponseWriter, r *http.Request) bool {
	if !strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "/" {
		return false
	}

	http.Redirect(w, r, strings.TrimRight(r.URL.Path, "/"), http.StatusTemporaryRedirect)

	return true
}
