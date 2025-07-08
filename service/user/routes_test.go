package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/absakran01/ecom/types"
)

type mockUserStore struct{}

// CreateUser implements types.UserStore.
func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayLoad{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "", // invalid: empty email
			Password:  "password123",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

// getUserByID implements types.UserStore.
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
