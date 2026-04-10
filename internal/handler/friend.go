package handler

import (
	"co-viewing/internal/models"
	"encoding/json"
	"net/http"
)


// patch accept friend request
func (h *Handler) HandleCreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request models.FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := h.svc.CreateFriendRequest(ctx, request.UserId, request.FriendId); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "database error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"created": "ok"})
}

func 