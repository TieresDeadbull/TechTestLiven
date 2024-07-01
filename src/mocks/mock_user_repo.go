package mocks

import "api/src/models"

type MockUserRepository struct {
	GetUserByEmailFunc func(email string) (*models.User, error)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	return m.GetUserByEmailFunc(email)
}
