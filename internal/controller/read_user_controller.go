package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamada-ai/go-twitter/api"
	"github.com/yamada-ai/go-twitter/internal/usecase"
)

type Server struct {
	userUsecase *usecase.ReadUserUsecase
}

func NewServer(u *usecase.ReadUserUsecase) *Server {
	return &Server{userUsecase: u}
}

// GetUsersId implements api.ServerInterface
func (s *Server) GetUsersId(ctx echo.Context, id int) error {
	user, err := s.userUsecase.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	resp := api.UserResponse{
		Id:       user.ID.Value(),
		Username: user.Username,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostUsers(ctx echo.Context) error {
	// TODO: リクエストを受け取り、新規ユーザーを作成するロジックを実装
	return ctx.JSON(http.StatusCreated, api.UserResponse{
		Id:       123,
		Username: "newuser",
	})
}
