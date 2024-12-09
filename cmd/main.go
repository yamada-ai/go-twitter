package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yamada-ai/go-twitter/api"
	"github.com/yamada-ai/go-twitter/internal/controller"
	"github.com/yamada-ai/go-twitter/internal/domain"
	"github.com/yamada-ai/go-twitter/internal/infrastructure"
	"github.com/yamada-ai/go-twitter/internal/usecase"
)

func main() {
	e := echo.New()

	// モックリポジトリに1ユーザー追加
	mockRepo := infrastructure.NewMockUserRepository()
	user, err := domain.NewUser(1, "testuser", "password123")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	mockRepo.AddUser(user)

	userU := usecase.NewUserUsecase(mockRepo)
	server := controller.NewServer(userU)

	api.RegisterHandlers(e, server) // OpenAPIで生成されたルーターにハンドラ登録

	e.Logger.Fatal(e.Start(":8080"))
}
