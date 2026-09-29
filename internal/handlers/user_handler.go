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

// GetMasteryProgress godoc
// GET /api/v1/user/kosakata/progress
func (h *UserHandler) GetMasteryProgress(w http.ResponseWriter, r *http.Request) {
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

	progress, err := h.service.GetMasteryProgress(userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// We can return the progress directly as an array of UserKosakata models
	respondJSON(w, http.StatusOK, progress)
}

type UpdateMasteryRequest struct {
	KosakataID uint `json:"kosakata_id"`
	IsCorrect  bool `json:"is_correct"`
}

// UpdateMastery godoc
// POST /api/v1/user/kosakata/mastery
func (h *UserHandler) UpdateMastery(w http.ResponseWriter, r *http.Request) {
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

	var req UpdateMasteryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.UpdateMastery(userID, req.KosakataID, req.IsCorrect); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "mastery updated successfully"})
}

type AddExpRequest struct {
	CorrectAnswers int `json:"correct_answers"`
}

// AddExp godoc
// POST /api/v1/user/exp/add
func (h *UserHandler) AddExp(w http.ResponseWriter, r *http.Request) {
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

	var req AddExpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	expGained := req.CorrectAnswers * 10
	if expGained > 0 {
		detail, isLevelUp, err := h.service.AddExp(userID, expGained)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		respondJSON(w, http.StatusOK, map[string]any{
			"message": "exp added", 
			"exp_gained": expGained,
			"is_level_up": isLevelUp,
			"new_level": detail.Level,
		})
		return
	}

	respondJSON(w, http.StatusOK, map[string]any{"message": "exp added", "exp_gained": expGained})
}

// GetProfile godoc
// GET /api/v1/user/profile
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
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

	profile, err := h.service.GetProfile(userID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, profile)
}
