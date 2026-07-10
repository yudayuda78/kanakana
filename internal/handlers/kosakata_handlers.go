package handlers

import (
	"encoding/json"
	"kanakana/internal/models"
	"net/http"
)

func GetKosakata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(daftarKosakata)
}

func PostKosakata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	daftarKosakata = append(daftarKosakata, kosakata)

	json.NewEncoder(w).Encode(kosakata)
}

func UpdateKosakata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, v := range daftarKosakata {
		if v.ID == kosakata.ID {
			daftarKosakata[i] = kosakata
			json.NewEncoder(w).Encode(kosakata)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func DeleteKosakata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, v := range daftarKosakata {
		if v.ID == kosakata.ID {
			daftarKosakata = append(daftarKosakata[:i], daftarKosakata[i+1:]...)
			json.NewEncoder(w).Encode(kosakata)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func GetKosakataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range daftarKosakata {
		if v.ID == kosakata.ID {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func GetKosakataByLevelJLPT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range daftarKosakata {
		if v.LevelJLPT == kosakata.LevelJLPT {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func GetKosakataByKanji(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range daftarKosakata {
		if v.Kanji == kosakata.Kanji {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func GetKosakataByReading(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range daftarKosakata {
		if v.Reading == kosakata.Reading {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}

func GetKosakataByRomaji(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kosakata models.Kosakata
	if err := json.NewDecoder(r.Body).Decode(&kosakata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, v := range daftarKosakata {
		if v.Romaji == kosakata.Romaji {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Kosakata tidak ditemukan", http.StatusNotFound)
}
