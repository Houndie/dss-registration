package commontest

import (
	"context"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/square-go/catalog"
	"github.com/Houndie/square-go/checkout"
	"github.com/Houndie/square-go/inventory"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
)

type MockSquareCatalogClient struct {
	catalog.Client
	ListFunc func(ctx context.Context, types []objects.CatalogObjectType) catalog.ListIterator
}

type MockSquareInventoryClient struct {
	inventory.Client
	BatchRetrieveCountsFunc func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator
}

type MockSquareLocationsClient struct {
	locations.Client
	ListFunc func(ctx context.Context) ([]*objects.Location, error)
}

type MockSquareCheckoutClient struct {
	checkout.Client
	CreateFunc func(ctx context.Context, locationID, idempotencyKey string, order *objects.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *objects.Address, redirectUrl string, additionalRecipients []*objects.ChargeRequestAdditionalRecipient, note string) (*objects.Checkout, error)
}

type MockSquareOrdersClient struct {
	orders.Client
	BatchRetrieveFunc func(ctx context.Context, locationID string, orderIDs []string) ([]*objects.Order, error)
}

func (m *MockSquareCatalogClient) List(ctx context.Context, types []objects.CatalogObjectType) catalog.ListIterator {
	return m.ListFunc(ctx, types)
}

func (m *MockSquareInventoryClient) BatchRetrieveCounts(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator {
	return m.BatchRetrieveCountsFunc(ctx, catalogObjectIDs, locationIDs, updatedAfter)
}

func (m *MockSquareLocationsClient) List(ctx context.Context) ([]*objects.Location, error) {
	return m.ListFunc(ctx)
}

func (m *MockSquareCheckoutClient) Create(ctx context.Context, locationID, idempotencyKey string, order *objects.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *objects.Address, redirectUrl string, additionalRecipients []*objects.ChargeRequestAdditionalRecipient, note string) (*objects.Checkout, error) {
	return m.CreateFunc(ctx, locationID, idempotencyKey, order, askForShippingAddress, merchantSupportEmail, prePopulateBuyerEmail, prePopulateShippingAddress, redirectUrl, additionalRecipients, note)
}

type MockListCatalogIterator struct {
	ValueFunc func() *objects.CatalogObject
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockListCatalogIterator) Value() *objects.CatalogObject {
	return m.ValueFunc()
}

func (m *MockListCatalogIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockListCatalogIterator) Next() bool {
	return m.NextFunc()
}

type MockBatchRetrieveInventoryCountsIterator struct {
	ValueFunc func() *objects.InventoryCount
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockBatchRetrieveInventoryCountsIterator) Value() *objects.InventoryCount {
	return m.ValueFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Next() bool {
	return m.NextFunc()
}

func (m *MockSquareOrdersClient) BatchRetrieve(ctx context.Context, locationID string, orderIDs []string) ([]*objects.Order, error) {
	return m.BatchRetrieveFunc(ctx, locationID, orderIDs)
}

/** Some common func implementations */

func ListCatalogFuncFromSlice(catalogObjects []*objects.CatalogObject) func(ctx context.Context, types []objects.CatalogObjectType) catalog.ListIterator {
	catalogObjectsIdx := -1
	return func(ctx context.Context, types []objects.CatalogObjectType) catalog.ListIterator {
		return &MockListCatalogIterator{
			ValueFunc: func() *objects.CatalogObject {
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

func InventoryCountsFromSliceCheck(t *testing.T, expectedObjectIDs map[storage.WeekendPassTier]string, inventoryCounts []*objects.InventoryCount) func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator {
	inventoryCountsIdx := -1
	return func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator {

		for _, expectedID := range expectedObjectIDs {
			found := false
			for _, passedID := range catalogObjectIDs {
				if passedID == expectedID {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("expected to find catalog object id %s in call, but did not", expectedID)
			}
		}

		for _, foundID := range catalogObjectIDs {
			found := false
			for _, allowedID := range expectedObjectIDs {
				if foundID == allowedID {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("found unexpected catalog object id %s", foundID)
			}
		}
		return &MockBatchRetrieveInventoryCountsIterator{
			ValueFunc: func() *objects.InventoryCount {
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

func InventoryCountsFromSlice(inventoryCounts []*objects.InventoryCount) func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator {
	inventoryCountsIdx := -1
	return func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) inventory.BatchRetrieveCountsIterator {
		return &MockBatchRetrieveInventoryCountsIterator{
			ValueFunc: func() *objects.InventoryCount {
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

func OrdersFromSliceCheck(t *testing.T, expectedLocationID string, orders []*objects.Order) func(ctx context.Context, locationID string, orderIDs []string) ([]*objects.Order, error) {
	return func(ctx context.Context, locationID string, orderIDs []string) ([]*objects.Order, error) {
		if locationID != expectedLocationID {
			t.Fatalf("found unexpected location id %s, expected %s", locationID, expectedLocationID)
		}

		retOrders := make([]*objects.Order, len(orderIDs))
		for i, orderID := range orderIDs {
			found := false
			for _, order := range orders {
				if orderID != order.ID {
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
