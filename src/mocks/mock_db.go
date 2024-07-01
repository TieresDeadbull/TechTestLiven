package mocks

import "database/sql"

type MockDB struct {
	ConnectFunc func() (*sql.DB, error)
}

func (m *MockDB) Connect() (*sql.DB, error) {
	return m.ConnectFunc()
}
