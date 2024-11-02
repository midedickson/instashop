package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/internal/http/controllers"
	"github.com/midedickson/instashop/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	mockUserService := new(mocks.MockUserService)
	mockProductService := new(mocks.MockProductService)

	ctrl := controllers.NewController(mockUserService, mockProductService)
	handler := http.HandlerFunc(ctrl.CreateUser)

	t.Run("successfully create a new user", func(t *testing.T) {
		signupPayload := dto.UserAuthPayload{
			Email:    "test@example.com",
			Password: "test123",
		}

		user := &entity.User{
			ID:    uint(123),
			Email: signupPayload.Email,
		}

		mockUserService.On("CreateUser", signupPayload).Return(user, nil)

		body, _ := json.Marshal(signupPayload)
		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, "User created successfully", response["message"])

		mockUserService.AssertExpectations(t)
	})

	t.Run("invalid JSON payload", func(t *testing.T) {
		// Send invalid JSON
		invalidJSON := `{"email": "test@example.com", "password": "test123` // missing closing brace

		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer([]byte(invalidJSON)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, "Invalid Payload", response["message"])
	})

	t.Run("missing required fields", func(t *testing.T) {
		// Missing password
		signupPayload := dto.UserAuthPayload{
			Email: "test@example.com",
		}

		body, _ := json.Marshal(signupPayload)
		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response map[string]interface{}
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, "Invalid Payload, email and password is required", response["message"])
	})

	t.Run("empty request body", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer([]byte{}))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("service returns error", func(t *testing.T) {
		signupPayload := dto.UserAuthPayload{
			Email:    "tester@example.com",
			Password: "test1234",
		}

		expectedError := errors.New("database error")
		mockUserService.On("CreateUser", signupPayload).Return((*entity.User)(nil), expectedError)

		body, _ := json.Marshal(signupPayload)
		req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUserService.AssertExpectations(t)
	})

}
