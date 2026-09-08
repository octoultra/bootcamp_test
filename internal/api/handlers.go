package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (a *API) health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "ok")
}

type echoPayload struct {
	Message string `json:"message"`
}

func (a *API) echo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.Write(body)
		return
	}

	var payload echoPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
