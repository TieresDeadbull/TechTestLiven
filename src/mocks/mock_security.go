package mocks

type MockSecurity struct {
	HashFunc       func(pass string) ([]byte, error)
	VerifyPassFunc func(hashedPass string, pass string) error
}

func (m *MockSecurity) Hash(pass string) ([]byte, error) {
	return m.HashFunc(pass)
}

func (m *MockSecurity) VerifyPass(hashedPass string, pass string) error {
	return m.VerifyPassFunc(hashedPass, pass)
}
