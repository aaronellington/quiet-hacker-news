package forge

import (
	"net/http"
	"strings"
)

// HTTPStatic servers static files without directory listings
type HTTPStatic struct {
	FileSystem      http.FileSystem
	NotFoundHandler http.Handler
	fileServer      http.Handler
}

// ServerHTTP satisfies the http.Handler interface
func (static *HTTPStatic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if static.fileServer == nil {
		static.fileServer = http.FileServer(static.FileSystem)
	}

	requestedFileName := r.URL.Path

	requestingDirectory := strings.HasSuffix(requestedFileName, "/")
	if requestingDirectory {
		requestedFileName += "index.html"
	}

	if !static.fileExists(requestedFileName) {
		static.notFound(w, r)
		return
	}

	w.Header().Add(HeaderCacheControl, "no-cache")
	static.fileServer.ServeHTTP(w, r)
}

func (static *HTTPStatic) notFound(w http.ResponseWriter, r *http.Request) {
	if static.NotFoundHandler != nil {
		static.NotFoundHandler.ServeHTTP(w, r)
		return
	}

	notFoundHandler(w, r)
}

func (static *HTTPStatic) fileExists(path string) bool {
	file, err := static.FileSystem.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	if fileInfo.IsDir() {
		path += "/index.html"

		return static.fileExists(path)
	}

	return true
}
