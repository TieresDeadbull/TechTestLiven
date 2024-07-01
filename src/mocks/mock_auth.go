package mocks

type MockAuth struct {
	CreateTokenFunc func(userID uint64) (string, error)
}

func (m *MockAuth) CreateToken(userID uint64) (string, error) {
	return m.CreateTokenFunc(userID)
}
