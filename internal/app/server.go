package app

import (
	"co-viewing/internal/handler"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"status", wrapped.status,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter(h *handler.Handler) http.Handler {
	router := mux.NewRouter()

	router.Path("/api/co-viewing/users/signup").Methods("POST").HandlerFunc(h.HandleCreateUser)
	router.Path("/api/co-viewing/users/signin").Methods("POST").HandlerFunc(h.HandleSignIn)
	router.Path("/api/co-viewing/users").Methods("GET").HandlerFunc(h.HandleGetAllUsers)
	router.Path("/api/co-viewing/friends").Methods("POST").HandlerFunc(h.HandleCreateFriendRequest)
	router.Path("/api/co-viewing/friends/{id}").Methods("GET").HandlerFunc(h.HandleGetFriendRequests)
	router.Path("/api/co-viewing/friends/{id}/list").Methods("GET").HandlerFunc(h.HandleGetAllFriends)
	router.Path("/api/co-viewing/friends").Methods("PATCH").HandlerFunc(h.HandlePatchFriendRequest)
	router.Path("/api/co-viewing/friends").Methods("DELETE").HandlerFunc(h.HandleDeleteFriendRequest)
	router.Path("/api/co-viewing/movies").Methods("POST").HandlerFunc(h.HandleAddWatchedMovie)
	router.Path("/api/co-viewing/movies/{id}").Methods("GET").HandlerFunc(h.HandleGetAllWatchedMovies)

	return corsMiddleware(loggingMiddleware(router))
}
