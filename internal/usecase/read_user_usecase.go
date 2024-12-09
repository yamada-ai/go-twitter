package usecase

import (
	"github.com/yamada-ai/go-twitter/internal/domain"
)

type ReadUserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *ReadUserUsecase {
	return &ReadUserUsecase{userRepo: repo}
}

func (u *ReadUserUsecase) GetUserByID(id int) (domain.User, error) {
	uid, err := domain.NewID[domain.User](id)
	if err != nil {
		return domain.User{}, err
	}
	user, err := u.userRepo.FindByID(uid)
	if err != nil {
		if err == domain.ErrUserNotFound {
			return domain.User{}, err
		}
		return domain.User{}, err
	}
	return user, nil
}
