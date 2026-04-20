package handler

import (
	"co-viewing/internal/models"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) HandleAddWatchedMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var movie models.WatchedMovie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		slog.Error("error", "invalid json", err)
		return
	}

	if movie.UserId <= 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be positive integer")
		slog.Error("id must be positive integer")
		return
	}

	if err := h.svc.AddWatchedMovie(ctx, movie); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		slog.Error("database error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"created": "ok"})
}

func (h *Handler) HandleGetAllWatchedMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be positive integer")
		slog.Error("id must be positive integer")
		return
	}

	movies, err := h.svc.GetAllWatchedMovies(ctx, id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		slog.Error("error", "database error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
