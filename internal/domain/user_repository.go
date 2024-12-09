package domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByID(id ID[User]) (User, error)
}
