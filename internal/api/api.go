package api

import (
	"net/http"
)

type API struct{}

func Init() *API {
	return &API{}
}

func (a *API) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", a.health)
	mux.HandleFunc("POST /echo", a.echo)
}
