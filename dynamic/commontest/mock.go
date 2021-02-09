package commontest

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

type MockMailClient struct {
	SendSMTPEmailFunc func(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error)
}

func (m *MockMailClient) SendSMTPEmail(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error) {
	return m.SendSMTPEmailFunc(ctx, params)
}

type MockStore struct {
	AddRegistrationFunc        func(context.Context, *storage.Registration) (string, error)
	GetDiscountFunc            func(context.Context, string) (*storage.Discount, error)
	AddDiscountFunc            func(context.Context, *storage.Discount) error
	GetRegistrationFunc        func(context.Context, string) (*storage.Registration, error)
	GetRegistrationsByUserFunc func(context.Context, string) ([]*storage.Registration, error)
	IsAdminFunc                func(context.Context, string) (bool, error)
	UpdateRegistrationFunc     func(ctx context.Context, r *storage.Registration) error
}

func (m *MockStore) AddRegistration(ctx context.Context, registration *storage.Registration) (string, error) {
	return m.AddRegistrationFunc(ctx, registration)
}

func (m *MockStore) GetDiscount(ctx context.Context, code string) (*storage.Discount, error) {
	return m.GetDiscountFunc(ctx, code)
}

func (m *MockStore) AddDiscount(ctx context.Context, discount *storage.Discount) error {
	return m.AddDiscountFunc(ctx, discount)
}

func (m *MockStore) IsAdmin(ctx context.Context, userID string) (bool, error) {
	return m.IsAdminFunc(ctx, userID)
}

func (m *MockStore) GetRegistrationsByUser(ctx context.Context, userID string) ([]*storage.Registration, error) {
	return m.GetRegistrationsByUserFunc(ctx, userID)
}

func (m *MockStore) GetRegistration(ctx context.Context, userID string) (*storage.Registration, error) {
	return m.GetRegistrationFunc(ctx, userID)
}

func (m *MockStore) UpdateRegistration(ctx context.Context, r *storage.Registration) error {
	return m.UpdateRegistrationFunc(ctx, r)
}
