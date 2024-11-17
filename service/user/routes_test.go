package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest/types"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Fail if payload is not passed", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("unexpected Status Code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Fail if payload is not valid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Username: "some_name",
			Email:    "invalid",
			Password: "",
		}
		marshelled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshelled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("unexpected Status Code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Pass if payload is valid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Username: "Jerald",
			Email:    "jerald@gmail.com",
			Password: "Jerald@123",
		}
		marshelled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshelled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("unexpected Status Code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m mockUserStore) CreateUser(user *types.User) error {
	return nil
}
