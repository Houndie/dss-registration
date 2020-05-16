package registration

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mockMailClient struct {
	SendFunc func(*mail.SGMailV3) (*rest.Response, error)
}

func (m *mockMailClient) Send(email *mail.SGMailV3) (*rest.Response, error) {
	return m.SendFunc(email)
}

type mockStore struct {
	AddDiscountFunc            func(context.Context, *storage.Discount) error
	AddRegistrationFunc        func(context.Context, *storage.Registration) (string, error)
	GetDiscountFunc            func(context.Context, string) (*storage.Discount, error)
	GetRegistrationFunc        func(context.Context, string) (*storage.Registration, error)
	GetRegistrationsByUserFunc func(context.Context, string) ([]*storage.Registration, error)
	IsAdminFunc                func(context.Context, string) (bool, error)
	UpdateRegistrationFunc     func(ctx context.Context, r *storage.Registration) error
}

func (m *mockStore) AddRegistration(ctx context.Context, registration *storage.Registration) (string, error) {
	return m.AddRegistrationFunc(ctx, registration)
}

func (m *mockStore) GetDiscount(ctx context.Context, code string) (*storage.Discount, error) {
	return m.GetDiscountFunc(ctx, code)
}

func (m *mockStore) AddDiscount(ctx context.Context, discount *storage.Discount) error {
	return m.AddDiscountFunc(ctx, discount)
}

func (m *mockStore) IsAdmin(ctx context.Context, userID string) (bool, error) {
	return m.IsAdminFunc(ctx, userID)
}

func (m *mockStore) GetRegistrationsByUser(ctx context.Context, userID string) ([]*storage.Registration, error) {
	return m.GetRegistrationsByUserFunc(ctx, userID)
}

func (m *mockStore) GetRegistration(ctx context.Context, userID string) (*storage.Registration, error) {
	return m.GetRegistrationFunc(ctx, userID)
}

func (m *mockStore) UpdateRegistration(ctx context.Context, r *storage.Registration) error {
	return m.UpdateRegistrationFunc(ctx, r)
}
