package domain

import "errors"

type User struct {
	ID       int
	Username string
	Password string
}

// NewUser: Userエンティティを作成するファクトリメソッド
func NewUser(id int, username, password string) (User, error) {
	if len(username) < 3 {
		return User{}, errors.New("username must be at least 3 characters long")
	}
	if len(password) < 6 {
		return User{}, errors.New("password must be at least 6 characters long")
	}
	return User{
		ID:       id,
		Username: username,
		Password: password,
	}, nil
}
