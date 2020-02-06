package server

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

const message = "hello world"

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	// its faster to provaid certain information
	// http status, content type ...
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte(message))
}

func InitHandler(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
