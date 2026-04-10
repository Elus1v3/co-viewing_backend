package handler

import (
	"co-viewing/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// patch accept friend request
func (h *Handler) HandleCreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request models.FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if request.UserId <= 0 || request.FriendId <= 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be positive integer")
	}

	if err := h.svc.CreateFriendRequest(ctx, request.UserId, request.FriendId); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"created": "ok"})
}

func (h *Handler) HandleGetFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be positive integer")
		return
	}

	users, err := h.svc.GetFriendRequestsFromId(ctx, id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) HandlePatchFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request models.FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if request.UserId <= 0 || request.FriendId <= 0 {
		writeJSONError(w, http.StatusBadRequest, "id must be positive integer")
		return
	}

	if err := h.svc.UpdateFriendRequest(ctx, request); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"updated": "ok"})
}
