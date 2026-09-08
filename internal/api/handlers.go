package api

import (
	"io"
	"net/http"
)

func (a *API) health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "ok")
}

func (a *API) echo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	contentType := r.Header.Get("Content-Type")
	body, err := io.ReadAll(r.Body)

	if err != nil || contentType != "text/plain" {
		http.Error(w, "type of content is not supported", http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
