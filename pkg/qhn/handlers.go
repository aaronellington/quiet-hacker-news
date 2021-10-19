package qhn

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/fuzzingbits/forge"
)

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyBuffer := bytes.NewBuffer([]byte{})
	if err := app.indexTemplate.Execute(bodyBuffer, app.itemCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := ioutil.ReadAll(bodyBuffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	forge.RespondHTML(w, http.StatusOK, responseBytes)
}
