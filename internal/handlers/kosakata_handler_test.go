package handlers

import (
	"bytes"
	"encoding/json"
	"kanakana/internal/models"
	"kanakana/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockKosakataService adalah implementasi mock dari services.KosakataService
type MockKosakataService struct {
	GetAllFunc          func() ([]models.Kosakata, error)
	GetByIDFunc         func(id uint) (*models.Kosakata, error)
	GetByKanjiFunc      func(kanji string) ([]models.Kosakata, error)
	GetByReadingFunc    func(reading string) ([]models.Kosakata, error)
	GetByRomajiFunc     func(romaji string) ([]models.Kosakata, error)
	GetByLevelJLPTFunc  func(level string) ([]models.Kosakata, error)
	CreateFunc          func(k *models.Kosakata) error
	UpdateFunc          func(id uint, k *models.Kosakata) (*models.Kosakata, error)
	DeleteFunc          func(id uint) error
}

func (m *MockKosakataService) GetAll() ([]models.Kosakata, error) {
	return m.GetAllFunc()
}

func (m *MockKosakataService) GetByID(id uint) (*models.Kosakata, error) {
	return m.GetByIDFunc(id)
}

func (m *MockKosakataService) GetByKanji(kanji string) ([]models.Kosakata, error) {
	return m.GetByKanjiFunc(kanji)
}

func (m *MockKosakataService) GetByReading(reading string) ([]models.Kosakata, error) {
	return m.GetByReadingFunc(reading)
}

func (m *MockKosakataService) GetByRomaji(romaji string) ([]models.Kosakata, error) {
	return m.GetByRomajiFunc(romaji)
}

func (m *MockKosakataService) GetByLevelJLPT(level string) ([]models.Kosakata, error) {
	return m.GetByLevelJLPTFunc(level)
}

func (m *MockKosakataService) Create(k *models.Kosakata) error {
	return m.CreateFunc(k)
}

func (m *MockKosakataService) Update(id uint, k *models.Kosakata) (*models.Kosakata, error) {
	return m.UpdateFunc(id, k)
}

func (m *MockKosakataService) Delete(id uint) error {
	return m.DeleteFunc(id)
}

func TestGetAll(t *testing.T) {
	mockService := &MockKosakataService{
		GetAllFunc: func() ([]models.Kosakata, error) {
			return []models.Kosakata{
				{Kanji: "日本語", Reading: "にほんご", Romaji: "Nihongo", Arti: "Bahasa Jepang", Levels: []models.KosakataLevel{{Name: "N5"}}},
			}, nil
		},
	}

	handler := NewKosakataHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/kosakata", nil)
	rr := httptest.NewRecorder()

	handler.GetAll(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusOK)
	}

	var response []models.Kosakata
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Gagal decode response body: %v", err)
	}

	if len(response) != 1 {
		t.Errorf("Ekspektasi 1 item, dapat %d", len(response))
	}
	if response[0].Kanji != "日本語" {
		t.Errorf("Ekspektasi kanji '日本語', dapat '%s'", response[0].Kanji)
	}
}

func TestGetByID_Success(t *testing.T) {
	mockService := &MockKosakataService{
		GetByIDFunc: func(id uint) (*models.Kosakata, error) {
			return &models.Kosakata{Kanji: "食べる", Levels: []models.KosakataLevel{{Name: "N5"}}}, nil
		},
	}

	handler := NewKosakataHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/kosakata/detail?id=1", nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusOK)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	mockService := &MockKosakataService{
		GetByIDFunc: func(id uint) (*models.Kosakata, error) {
			return nil, services.ErrNotFound
		},
	}

	handler := NewKosakataHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/kosakata/detail?id=99", nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusNotFound)
	}
}

func TestCreate_Success(t *testing.T) {
	mockService := &MockKosakataService{
		CreateFunc: func(k *models.Kosakata) error {
			return nil
		},
	}

	handler := NewKosakataHandler(mockService)

	newKosaKata := models.Kosakata{
		Kanji: "先生", Reading: "せんせい", Romaji: "Sensei", Arti: "Guru", Levels: []models.KosakataLevel{{Name: "N5"}},
	}
	body, _ := json.Marshal(newKosaKata)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/kosakata", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusCreated)
	}
}

func TestUpdate_Success(t *testing.T) {
	mockService := &MockKosakataService{
		UpdateFunc: func(id uint, k *models.Kosakata) (*models.Kosakata, error) {
			return k, nil
		},
	}

	handler := NewKosakataHandler(mockService)

	updateKosaKata := models.Kosakata{
		Kanji: "先生", Reading: "せんせい", Romaji: "Sensei", Arti: "Guru", Levels: []models.KosakataLevel{{Name: "N5"}},
	}
	body, _ := json.Marshal(updateKosaKata)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/kosakata?id=1", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusOK)
	}
}

func TestDelete_Success(t *testing.T) {
	mockService := &MockKosakataService{
		DeleteFunc: func(id uint) error {
			return nil
		},
	}

	handler := NewKosakataHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/kosakata?id=1", nil)
	rr := httptest.NewRecorder()

	handler.Delete(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return status salah: dapat %v ingin %v", status, http.StatusOK)
	}
}
