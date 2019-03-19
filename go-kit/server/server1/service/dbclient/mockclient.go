package dbclient

import (
	"github.com/callistaenterprise/goblog/accountservice/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock implementation of a datastore client for testing purposes.
// Instead of the bolt.DB pointer, we're just putting a generic mock object from
// strechr/testify
type MockBoltClient struct {
	mock.Mock
}

// From here, we'll declare three functions that makes our MockBoltClient fulfill the interface IBoltClient that we declared in part 3.
func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}
