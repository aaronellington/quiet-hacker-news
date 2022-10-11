package forge

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// HTTPStatic is foobar
type HTTPStatic struct {
	FileSystem      http.FileSystem
	NotFoundHandler http.Handler
}

// ServeHTTP is foobar
func (httpStatic *HTTPStatic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedFileName := r.URL.Path
	isRequestingDirectory := strings.HasSuffix(requestedFileName, "/")
	if isRequestingDirectory {
		requestedFileName += "index.html"
	}

	file, err := httpStatic.FileSystem.Open(requestedFileName)
	if err != nil {
		correctNotFoundHandler(httpStatic.NotFoundHandler).ServeHTTP(w, r)
		return
	}
	defer file.Close()
	fileTypeHeader := mime.TypeByExtension(filepath.Ext(requestedFileName))

	w.Header().Set("Content-Type", fileTypeHeader)
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}
