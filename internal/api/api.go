package api

import (
	"entrytest/internal/store"
	"net/http"
)

type API struct {
	store *store.MessageStore
}

func Init(store *store.MessageStore) *API {
	return &API{store: store}
}

func (a *API) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", a.health)
	mux.HandleFunc("POST /echo", a.echo)
	mux.HandleFunc("POST /messages", a.createMessage)
	mux.HandleFunc("GET /messages", a.listMessages)
}
