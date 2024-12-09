package infrastructure

import "github.com/yamada-ai/go-twitter/internal/domain"

type MockUserRepository struct {
	users map[int]domain.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: map[int]domain.User{},
	}
}

func (m *MockUserRepository) AddUser(u domain.User) {
	m.users[u.ID.Value()] = u
}

func (m *MockUserRepository) FindByID(id domain.ID[domain.User]) (domain.User, error) {
	user, ok := m.users[id.Value()]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}
