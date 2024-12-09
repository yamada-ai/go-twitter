package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yamada-ai/go-twitter/api"
	"github.com/yamada-ai/go-twitter/internal/domain"
	"github.com/yamada-ai/go-twitter/internal/infrastructure"
	"github.com/yamada-ai/go-twitter/internal/usecase"
)

func TestGetUserByID(t *testing.T) {
	e := echo.New()
	mockRepo := infrastructure.NewMockUserRepository()

	user, _ := domain.NewUser(1, "testuser", "password123")
	mockRepo.AddUser(user)

	userU := usecase.NewUserUsecase(mockRepo)
	server := NewServer(userU)
	api.RegisterHandlers(e, server)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp api.UserResponse
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.Id)
	assert.Equal(t, "testuser", resp.Username)
}
