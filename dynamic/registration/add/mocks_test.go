package add

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/square"
)

type mockListCatalogIterator struct {
	ValueFunc func() *square.CatalogObject
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *mockListCatalogIterator) Value() *square.CatalogObject {
	return m.ValueFunc()
}

func (m *mockListCatalogIterator) Error() error {
	return m.ErrorFunc()
}

func (m *mockListCatalogIterator) Next() bool {
	return m.NextFunc()
}

type mockSquareClient struct {
	ListCatalogFunc    func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	ListLocationsFunc  func(ctx context.Context) ([]*square.Location, error)
	CreateCheckoutFunc func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
}

func (m *mockSquareClient) ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
	return m.ListCatalogFunc(ctx, types)
}

func (m *mockSquareClient) ListLocations(ctx context.Context) ([]*square.Location, error) {
	return m.ListLocationsFunc(ctx)
}

func (m *mockSquareClient) CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
	return m.CreateCheckoutFunc(ctx, locationId, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)
}

type mockStore struct {
	AddRegistrationFunc    func(context.Context, *StoreRegistration) (string, error)
	DeleteRegistrationFunc func(context.Context, string) error
}

func (m *mockStore) AddRegistration(ctx context.Context, r *StoreRegistration) (string, error) {
	return m.AddRegistrationFunc(ctx, r)
}

func (m *mockStore) DeleteRegistration(ctx context.Context, s string) error {
	return m.DeleteRegistrationFunc(ctx, s)
}
