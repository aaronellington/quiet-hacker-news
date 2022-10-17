package forge

import "net/http"

func correctNotFoundHandler(customHandler http.Handler) http.Handler {
	if customHandler != nil {
		return customHandler
	}

	return http.NotFoundHandler()
}

// UnreachableError is foobar
func UnreachableError(err error) {
	if err != nil {
		panic(err)
	}
}