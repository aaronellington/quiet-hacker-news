package forge

import (
	"encoding/json"
	"net/http"
)

// RespondJSON is foobar
func RespondJSON(w http.ResponseWriter, status int, v interface{}) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(status)
	encoder.Encode(v)
}

// RespondHTML is foobar
func RespondHTML(w http.ResponseWriter, status int, s string) {
	w.WriteHeader(status)
	w.Write([]byte(s))
}
