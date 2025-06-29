package routes_test

import (
	"bytes"
	"capsuler/pkg/capsule"
	"capsuler/pkg/routes"
	"capsuler/pkg/testify"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockCapsuleRepository struct {
	DB map[string]*capsule.Capsule
}

func NewMockCapsuleRepository() *mockCapsuleRepository {
	return &mockCapsuleRepository{
		DB: make(map[string]*capsule.Capsule),
	}
}

func (m *mockCapsuleRepository) Save(capsule *capsule.Capsule) error {
	m.DB[capsule.GetId()] = capsule
	return nil
}

func (m *mockCapsuleRepository) GetById(id string) (*capsule.Capsule, error) {
	return m.DB[id], nil
}

func TestCapsuleRoutes(t *testing.T) {
	mockCapsuleRepository := NewMockCapsuleRepository()

	t.Run("should return an 201 created with capsule id in location in headers", func(t *testing.T) {
		// Arrange
		reqBody := map[string]any{
			"name":         "capsule_test",
			"description":  "capsule test 2025",
			"date_to_open": time.Now().Unix(),
		}
		jsonBody, _ := json.Marshal(reqBody)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/capsules", bytes.NewBuffer(jsonBody))
		handler := routes.CreateCapsule(mockCapsuleRepository)

		// Act
		handler(w, r)

		// Assert
		testify.AssertEqualsInt(t, http.StatusCreated, w.Code)
		testify.AssertNotEmptyStr(t, w.Header().Get("Location"))
	})

	t.Run("should return an 200 ok with capsule id to open the capsule", func(t *testing.T) {
		// Arrange
		capsule := capsule.Builder().Build()
		mockCapsuleRepository.DB[capsule.GetId()] = capsule

		mux := http.NewServeMux()
		mux.HandleFunc("/capsules/{id}/open", routes.OpenCapsule(mockCapsuleRepository))

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", fmt.Sprintf("/capsules/%s/open", capsule.GetId()), nil)

		// Act
		mux.ServeHTTP(w, r)

		// Assert
		testify.AssertEqualsInt(t, http.StatusOK, w.Code)
		testify.AssertTrue(t, capsule.WasOpened())
	})

	t.Run("should return an 201 ok with capsule id to add a new message", func(t *testing.T) {
		// Arrange
		reqBody := map[string]any{
			"message": "message test",
		}
		jsonBody, _ := json.Marshal(reqBody)
		capsule := capsule.Builder().Build()
		mockCapsuleRepository.DB[capsule.GetId()] = capsule

		mux := http.NewServeMux()
		mux.HandleFunc("/capsules/{id}/messages", routes.AddMessageToCapsule(mockCapsuleRepository))

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", fmt.Sprintf("/capsules/%s/messages", capsule.GetId()), bytes.NewBuffer(jsonBody))

		// Act
		mux.ServeHTTP(w, r)

		var mapData map[string]any
		testify.AssertNil(t, json.NewDecoder(w.Body).Decode(&mapData))

		messages := mapData["messages"].([]any)

		// Assert
		testify.AssertEqualsInt(t, http.StatusCreated, w.Code)
		testify.AssertNotEmptyAny(t, messages)

		for _, message := range messages {
			messageStr := message.(string)
			testify.AssertNotEmptyStr(t, messageStr)
		}
	})
}
