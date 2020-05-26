package commontest

import (
	"context"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

type MockSquareClient struct {
	ListCatalogFunc                  func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	BatchRetrieveInventoryCountsFunc func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator
	ListLocationsFunc                func(ctx context.Context) ([]*square.Location, error)
	CreateCheckoutFunc               func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
	BatchRetrieveOrdersFunc          func(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error)
}

func (m *MockSquareClient) ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
	return m.ListCatalogFunc(ctx, types)
}

func (m *MockSquareClient) BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
	return m.BatchRetrieveInventoryCountsFunc(ctx, catalogObjectIds, locationIds, updatedAfter)
}

func (m *MockSquareClient) ListLocations(ctx context.Context) ([]*square.Location, error) {
	return m.ListLocationsFunc(ctx)
}

func (m *MockSquareClient) CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
	return m.CreateCheckoutFunc(ctx, locationId, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)
}

type MockListCatalogIterator struct {
	ValueFunc func() *square.CatalogObject
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockListCatalogIterator) Value() *square.CatalogObject {
	return m.ValueFunc()
}

func (m *MockListCatalogIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockListCatalogIterator) Next() bool {
	return m.NextFunc()
}

type MockBatchRetrieveInventoryCountsIterator struct {
	ValueFunc func() *square.InventoryCount
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockBatchRetrieveInventoryCountsIterator) Value() *square.InventoryCount {
	return m.ValueFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Next() bool {
	return m.NextFunc()
}

func (m *MockSquareClient) BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error) {
	return m.BatchRetrieveOrdersFunc(ctx, locationId, orderIds)
}

/** Some common func implementations */

func ListCatalogFuncFromSlice(catalogObjects []*square.CatalogObject) func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
	catalogObjectsIdx := -1
	return func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
		return &MockListCatalogIterator{
			ValueFunc: func() *square.CatalogObject {
				return catalogObjects[catalogObjectsIdx]
			},
			ErrorFunc: func() error {
				return nil
			},
			NextFunc: func() bool {
				catalogObjectsIdx++
				return catalogObjectsIdx < len(catalogObjects)
			},
		}
	}
}

func InventoryCountsFromSliceCheck(t *testing.T, expectedObjectIds map[storage.WeekendPassTier]string, inventoryCounts []*square.InventoryCount) func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
	inventoryCountsIdx := -1
	return func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {

		for _, expectedId := range expectedObjectIds {
			found := false
			for _, passedId := range catalogObjectIds {
				if passedId == expectedId {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("expected to find catalog object id %s in call, but did not", expectedId)
			}
		}

		for _, foundId := range catalogObjectIds {
			found := false
			for _, allowedId := range expectedObjectIds {
				if foundId == allowedId {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("found unexpected catalog object id %s", foundId)
			}
		}
		return &MockBatchRetrieveInventoryCountsIterator{
			ValueFunc: func() *square.InventoryCount {
				return inventoryCounts[inventoryCountsIdx]
			},
			ErrorFunc: func() error {
				return nil
			},
			NextFunc: func() bool {
				inventoryCountsIdx++
				return inventoryCountsIdx < len(inventoryCounts)
			},
		}
	}
}

func InventoryCountsFromSlice(inventoryCounts []*square.InventoryCount) func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
	inventoryCountsIdx := -1
	return func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
		return &MockBatchRetrieveInventoryCountsIterator{
			ValueFunc: func() *square.InventoryCount {
				return inventoryCounts[inventoryCountsIdx]
			},
			ErrorFunc: func() error {
				return nil
			},
			NextFunc: func() bool {
				inventoryCountsIdx++
				return inventoryCountsIdx < len(inventoryCounts)
			},
		}
	}
}

func OrdersFromSliceCheck(t *testing.T, expectedLocationID string, orders []*square.Order) func(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error) {
	return func(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error) {
		if locationId != expectedLocationID {
			t.Fatalf("found unexpected location id %s, expected %s", locationId, expectedLocationID)
		}

		retOrders := make([]*square.Order, len(orderIds))
		for i, orderID := range orderIds {
			found := false
			for _, order := range orders {
				if orderID != order.Id {
					continue
				}
				found = true
				retOrders[i] = order
				break
			}
			if !found {
				t.Fatalf("unable to find order id %s in order id list", orderID)
			}
		}
		return orders, nil
	}
}
