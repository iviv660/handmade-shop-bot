package tests

import (
	"testing"

	"app/internal/service"
	"app/internal/service/tests/mocks"
)

type TestDeps struct {
	UC       *service.UseCase
	UserMock *mocks.MockUser
	ProdMock *mocks.MockProduct
	OrdMock  *mocks.MockOrder
	PayMock  *mocks.MockPayment
	LogMock  *mocks.MockLogger
}

func NewTestDeps(t *testing.T) *TestDeps {
	t.Helper()

	us := mocks.NewMockUser(t)
	ps := mocks.NewMockProduct(t)
	os := mocks.NewMockOrder(t)
	pays := mocks.NewMockPayment(t)
	log := mocks.NewMockLogger(t)

	uc := service.NewUseCase(us, ps, os, pays, log)

	t.Cleanup(func() {
		us.AssertExpectations(t)
		ps.AssertExpectations(t)
		os.AssertExpectations(t)
		pays.AssertExpectations(t)
		log.AssertExpectations(t)
	})

	return &TestDeps{
		UC:       uc,
		UserMock: us,
		ProdMock: ps,
		OrdMock:  os,
		PayMock:  pays,
		LogMock:  log,
	}
}
