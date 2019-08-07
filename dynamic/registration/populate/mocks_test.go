package populate

import (
	"context"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
)

type mockSquareClient struct {
	ListCatalogFunc                  func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	BatchRetrieveInventoryCountsFunc func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator
}

func (m *mockSquareClient) ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
	return m.ListCatalogFunc(ctx, types)
}

func (m *mockSquareClient) BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
	return m.BatchRetrieveInventoryCountsFunc(ctx, catalogObjectIds, locationIds, updatedAfter)
}

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

type mockBatchRetrieveInventoryCountsIterator struct {
	ValueFunc func() *square.InventoryCount
	ErrorFunc func() error
	NextFunc  func() bool
}

func (m *mockBatchRetrieveInventoryCountsIterator) Value() *square.InventoryCount {
	return m.ValueFunc()
}

func (m *mockBatchRetrieveInventoryCountsIterator) Error() error {
	return m.ErrorFunc()
}

func (m *mockBatchRetrieveInventoryCountsIterator) Next() bool {
	return m.NextFunc()
}
