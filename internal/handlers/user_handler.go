package handlers

import (
	"encoding/json"
	"kanakana/internal/middleware"
	"kanakana/internal/services"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(svc services.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

// GetMyLevels godoc
// GET /api/v1/user/levels
func (h *UserHandler) GetMyLevels(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value(middleware.UserIDKey)
	if userIDRaw == nil {
		respondError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	
	// Type assertion based on how you store it in your middleware (usually float64 if parsed from JWT)
	var userID uint
	switch v := userIDRaw.(type) {
	case float64:
		userID = uint(v)
	case uint:
		userID = v
	default:
		respondError(w, http.StatusInternalServerError, "invalid user id type in context")
		return
	}

	levels, err := h.service.GetUserLevels(userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, levels)
}

type UpdateMyLevelsRequest struct {
	LevelIDs []uint `json:"level_ids"`
}

// UpdateMyLevels godoc
// PUT /api/v1/user/levels
func (h *UserHandler) UpdateMyLevels(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value(middleware.UserIDKey)
	if userIDRaw == nil {
		respondError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	
	var userID uint
	switch v := userIDRaw.(type) {
	case float64:
		userID = uint(v)
	case uint:
		userID = v
	default:
		respondError(w, http.StatusInternalServerError, "invalid user id type in context")
		return
	}

	var req UpdateMyLevelsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateUserLevels(userID, req.LevelIDs); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "levels updated successfully"})
}
