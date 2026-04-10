package handler

import (
	"co-viewing/internal/models"
	"encoding/json"
	"net/http"
)

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if len(user.Nickname) > 20 || user.Nickname == "" {
		writeJSONError(w, http.StatusBadRequest, "nickname not contain over 20 symbols or empty string")
		return
	}

	if len(user.Password) > 255 || user.Password == "" {
		writeJSONError(w, http.StatusBadRequest, "password not contain over 255 symbols or empty string")
		return
	}

	id, err := h.svc.Create(ctx, user)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
	}

	if len(user.Nickname) > 20 || user.Nickname == "" {
		writeJSONError(w, http.StatusBadRequest, "nickname not contain over 20 symbols or empty string")
		return
	}

	if len(user.Password) > 255 || user.Password == "" {
		writeJSONError(w, http.StatusBadRequest, "password not contain over 255 symbols or empty string")
		return
	}

	user, err := h.svc.SignIn(ctx, user)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"id": user.Id})
}
