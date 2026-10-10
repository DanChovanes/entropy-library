// Package server provides an HTTP API for the entropy library.
package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"entropy-library"
)

const maxBytes = 1024

type entropyResponse struct {
	Data   []byte `json:"data"`
	Length int    `json:"length"`
}

// NewHandler returns an HTTP handler with health and entropy endpoints.
// GET /healthz reports server health.
// GET /v1/entropy?bytes=16 returns between 1 and 1024 random bytes.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /v1/entropy", handleEntropy)
	return mux
}

func handleEntropy(w http.ResponseWriter, r *http.Request) {
	length := 16
	if value := r.URL.Query().Get("bytes"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 1 || parsed > maxBytes {
			http.Error(w, "`bytes` must be an integer between 1 and 1024", http.StatusBadRequest)
			return
		}
		length = parsed
	}

	data, err := entropy.Bytes(length)
	if err != nil {
		http.Error(w, "failed to generate entropy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entropyResponse{Data: data, Length: len(data)}); err != nil {
		return
	}
}
