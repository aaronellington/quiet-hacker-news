package forge

import (
	"encoding/json"
	"net/http"
)

// Header Constants
const (
	HeaderContentType  = "Content-Type"
	HeaderCacheControl = "Cache-Control"
)

// Response Constants
const (
	ResponseTextNotFound = "Not Found"
)

// Response is a basic response structure
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespondText responds to an http.Request with a text body
func RespondText(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

// RespondHTML responds to an http.Request with a text body
func RespondHTML(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set(HeaderContentType, "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

// RespondJSON responds to an http.Request with a JSON body
func RespondJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set(HeaderContentType, "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(v)
}
