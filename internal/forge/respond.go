package forge

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, v interface{}) {
	encoder := json.NewEncoder(w)
	w.WriteHeader(status)
	encoder.Encode(v)
}

func RespondHTML(w http.ResponseWriter, status int, s string) {
	w.WriteHeader(status)
	w.Write([]byte(s))
}
