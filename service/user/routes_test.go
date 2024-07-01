package user

import (
	"bytes"
	"ecommerce-project/types"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore()
	handler := NewHandler(userStore)
	t.Run("shoud fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Berik",
			LastName:  "Yerman",
			Email:     "",
			Password:  "asd",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/regsiter", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d,got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
