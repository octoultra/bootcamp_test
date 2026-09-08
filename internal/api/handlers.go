package api

import (
	"io"
	"net/http"
)

func (a *API) health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "ok")
}
