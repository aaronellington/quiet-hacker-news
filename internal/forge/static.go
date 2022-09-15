package forge

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// Static is foobar
type Static struct {
	FileSystem      http.FileSystem
	NotFoundHandler http.Handler
}

// ServeHTTP is foobar
func (static *Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedFileName := r.URL.Path
	isRequestingDirectory := strings.HasSuffix(requestedFileName, "/")
	if isRequestingDirectory {
		requestedFileName += "index.html"
	}

	file, err := static.FileSystem.Open(requestedFileName)
	if err != nil {
		correctNotFoundHandler(static.NotFoundHandler).ServeHTTP(w, r)
		return
	}
	defer file.Close()

	fileTypeHeader := mime.TypeByExtension(filepath.Ext(requestedFileName))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", fileTypeHeader)
	io.Copy(w, file)
}
