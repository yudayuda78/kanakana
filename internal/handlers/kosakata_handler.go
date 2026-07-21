package handlers

import (
	"encoding/json"
	"errors"
	"kanakana/internal/models"
	"kanakana/internal/services"
	"net/http"
	"strconv"
)

// KosakataHandler menangani HTTP request untuk endpoint kosakata.
// Dependency diinjeksikan melalui constructor agar mudah di-test.
type KosakataHandler struct {
	service services.KosakataService
}

// NewKosakataHandler membuat instance baru KosakataHandler.
func NewKosakataHandler(svc services.KosakataService) *KosakataHandler {
	return &KosakataHandler{service: svc}
}

// GetAll godoc
// GET /api/v1/kosakata
func (h *KosakataHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.GetAll()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// GetAllLevels godoc
// GET /api/v1/levels
func (h *KosakataHandler) GetAllLevels(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.GetAllLevels()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// GetByID godoc
// GET /api/v1/kosakata/{id}
func (h *KosakataHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromQuery(r, "id")
	if err != nil {
		respondError(w, http.StatusBadRequest, "parameter 'id' tidak valid")
		return
	}

	k, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, k)
}

// GetByLevelJLPT godoc
// GET /api/v1/kosakata/level?level=N5
func (h *KosakataHandler) GetByLevelJLPT(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")
	if level == "" {
		respondError(w, http.StatusBadRequest, "parameter 'level' tidak boleh kosong")
		return
	}

	list, err := h.service.GetByLevelJLPT(level)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// GetByKanji godoc
// GET /api/v1/kosakata/kanji?q=日本語
func (h *KosakataHandler) GetByKanji(w http.ResponseWriter, r *http.Request) {
	kanji := r.URL.Query().Get("q")
	if kanji == "" {
		respondError(w, http.StatusBadRequest, "parameter 'q' tidak boleh kosong")
		return
	}

	list, err := h.service.GetByKanji(kanji)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// GetByReading godoc
// GET /api/v1/kosakata/reading?q=にほんご
func (h *KosakataHandler) GetByReading(w http.ResponseWriter, r *http.Request) {
	reading := r.URL.Query().Get("q")
	if reading == "" {
		respondError(w, http.StatusBadRequest, "parameter 'q' tidak boleh kosong")
		return
	}

	list, err := h.service.GetByReading(reading)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// GetByRomaji godoc
// GET /api/v1/kosakata/romaji?q=nihongo
func (h *KosakataHandler) GetByRomaji(w http.ResponseWriter, r *http.Request) {
	romaji := r.URL.Query().Get("q")
	if romaji == "" {
		respondError(w, http.StatusBadRequest, "parameter 'q' tidak boleh kosong")
		return
	}

	list, err := h.service.GetByRomaji(romaji)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

// Create godoc
// POST /api/v1/kosakata
func (h *KosakataHandler) Create(w http.ResponseWriter, r *http.Request) {
	var k models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		respondError(w, http.StatusBadRequest, "body request tidak valid")
		return
	}

	if err := h.service.Create(&k); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, k)
}

// Update godoc
// PUT /api/v1/kosakata/{id}
func (h *KosakataHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromQuery(r, "id")
	if err != nil {
		respondError(w, http.StatusBadRequest, "parameter 'id' tidak valid")
		return
	}

	var k models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		respondError(w, http.StatusBadRequest, "body request tidak valid")
		return
	}

	updated, err := h.service.Update(id, &k)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, updated)
}

// Delete godoc
// DELETE /api/v1/kosakata/{id}
func (h *KosakataHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromQuery(r, "id")
	if err != nil {
		respondError(w, http.StatusBadRequest, "parameter 'id' tidak valid")
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, services.ErrNotFound) {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "kosakata berhasil dihapus"})
}

// --- Helpers ---

// respondJSON menulis response JSON dengan status code yang diberikan.
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError menulis response error dalam format JSON.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// parseIDFromQuery mengambil dan mem-parse parameter uint dari query string.
func parseIDFromQuery(r *http.Request, key string) (uint, error) {
	val := r.URL.Query().Get(key)
	parsed, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}
