package commontest

import (
	"context"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/square-go/catalog"
	"github.com/Houndie/square-go/checkout"
	"github.com/Houndie/square-go/inventory"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
)

type MockSquareInventoryClient struct {
	inventory.Client
	BatchRetrieveCountsFunc func(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error)
}

type MockSquareLocationsClient struct {
	locations.Client
	ListFunc func(ctx context.Context, req *locations.ListRequest) (*locations.ListResponse, error)
}

type MockSquareCheckoutClient struct {
	checkout.Client
	CreateFunc func(ctx context.Context, req *checkout.CreateRequest) (*checkout.CreateResponse, error)
}

type MockSquareOrdersClient struct {
	orders.Client
	BatchRetrieveFunc func(ctx context.Context, req *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error)
}

func (m *MockSquareInventoryClient) BatchRetrieveCounts(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error) {
	return m.BatchRetrieveCountsFunc(ctx, req)
}

func (m *MockSquareLocationsClient) List(ctx context.Context, req *locations.ListRequest) (*locations.ListResponse, error) {
	return m.ListFunc(ctx, req)
}

func (m *MockSquareCheckoutClient) Create(ctx context.Context, req *checkout.CreateRequest) (*checkout.CreateResponse, error) {
	return m.CreateFunc(ctx, req)
}

type MockListCatalogIterator struct {
	ValueFunc func() *catalog.ListIteratorValue
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockListCatalogIterator) Value() *catalog.ListIteratorValue {
	return m.ValueFunc()
}

func (m *MockListCatalogIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockListCatalogIterator) Next() bool {
	return m.NextFunc()
}

type MockBatchRetrieveInventoryCountsIterator struct {
	ValueFunc func() *inventory.BatchRetrieveCountsIteratorValue
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *MockBatchRetrieveInventoryCountsIterator) Value() *inventory.BatchRetrieveCountsIteratorValue {
	return m.ValueFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Error() error {
	return m.ErrorFunc()
}

func (m *MockBatchRetrieveInventoryCountsIterator) Next() bool {
	return m.NextFunc()
}

func (m *MockSquareOrdersClient) BatchRetrieve(ctx context.Context, req *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error) {
	return m.BatchRetrieveFunc(ctx, req)
}

/** Some common func implementations */

func InventoryCountsFromSliceCheck(t *testing.T, expectedObjectIDs map[storage.WeekendPassTier]string, inventoryCounts []*objects.InventoryCount) func(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error) {
	inventoryCountsIdx := -1
	return func(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error) {

		for _, expectedID := range expectedObjectIDs {
			found := false
			for _, passedID := range req.CatalogObjectIDs {
				if passedID == expectedID {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("expected to find catalog object id %s in call, but did not", expectedID)
			}
		}

		for _, foundID := range req.CatalogObjectIDs {
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
		return &inventory.BatchRetrieveCountsResponse{
			Counts: &MockBatchRetrieveInventoryCountsIterator{
				ValueFunc: func() *inventory.BatchRetrieveCountsIteratorValue {
					return &inventory.BatchRetrieveCountsIteratorValue{
						Count: inventoryCounts[inventoryCountsIdx],
					}
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					inventoryCountsIdx++
					return inventoryCountsIdx < len(inventoryCounts)
				},
			},
		}, nil
	}
}

func InventoryCountsFromSlice(inventoryCounts []*objects.InventoryCount) func(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error) {
	inventoryCountsIdx := -1
	return func(ctx context.Context, req *inventory.BatchRetrieveCountsRequest) (*inventory.BatchRetrieveCountsResponse, error) {
		return &inventory.BatchRetrieveCountsResponse{
			Counts: &MockBatchRetrieveInventoryCountsIterator{
				ValueFunc: func() *inventory.BatchRetrieveCountsIteratorValue {
					return &inventory.BatchRetrieveCountsIteratorValue{
						Count: inventoryCounts[inventoryCountsIdx],
					}
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					inventoryCountsIdx++
					return inventoryCountsIdx < len(inventoryCounts)
				},
			},
		}, nil
	}
}

func OrdersFromSliceCheck(t *testing.T, expectedLocationID string, orderObjects []*objects.Order) func(ctx context.Context, req *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error) {
	return func(ctx context.Context, req *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error) {
		if req.LocationID != expectedLocationID {
			t.Fatalf("found unexpected location id %s, expected %s", req.LocationID, expectedLocationID)
		}

		retOrders := make([]*objects.Order, len(req.OrderIDs))
		for i, orderID := range req.OrderIDs {
			found := false
			for _, order := range orderObjects {
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
		return &orders.BatchRetrieveResponse{
			Orders: orderObjects,
		}, nil
	}
}
