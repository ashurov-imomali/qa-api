package app

import (
	"github.com/ashurov-imomali/qa-api/internal/handlers"
	"net/http"
	"time"
)

func NewServer(addr string, h *handlers.Handler) *http.Server {
	mux := http.NewServeMux()
	h.RegisterRouters(mux)

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv
}
