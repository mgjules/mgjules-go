package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func renderJSON(w http.ResponseWriter, status int, v any) {
	if v == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(v)
	if err != nil {
		slog.Error("failed to create a valid response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "could not create a valid response`))
		return
	}

	w.WriteHeader(status)
	w.Write(b)
}

func renderHTML(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Encoding", "gzip")
	if len(b) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
