package populate

import (
	"context"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func TestPopulate(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.AddHook(&test_utility.ErrorHook{t})
	logger.SetLevel(logrus.TraceLevel)

	catalogObjectsIdx := -1
	inventoryCountsIdx := -1

	client := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
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
		},
		BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
			expectedObjectIds := []string{
				weekendPassTier1VariationId,
				weekendPassTier2VariationId,
				weekendPassTier3VariationId,
				weekendPassTier4VariationId,
				weekendPassTier5VariationId,
			}

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
			return &mockBatchRetrieveInventoryCountsIterator{
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
		},
	}

	formData, err := NewService(logger, client).Populate(context.Background())
	if err != nil {
		t.Fatalf("error populating form data: %v", err)
	}

	if formData.WeekendPassCost != weekendPassCost {
		t.Fatalf("found unexpected weekend pass cost %d, expected %d", formData.WeekendPassCost, weekendPassCost)
	}

	if formData.WeekendPassTier != 3 {
		t.Fatalf("found unexpected weekend pass tier %d, expected 3", formData.WeekendPassTier)
	}

	if formData.DancePassCost != dancePassCost {
		t.Fatalf("found unexpected dance pass cost %d, expected %d", formData.DancePassCost, dancePassCost)
	}

	if formData.MixAndMatchCost != mixAndMatchCost {
		t.Fatalf("found unexpected mix and match cost %d, expected %d", formData.MixAndMatchCost, mixAndMatchCost)
	}

	if formData.SoloJazzCost != soloJazzCost {
		t.Fatalf("found unexpected solo jazz cost %d, expected %d", formData.SoloJazzCost, soloJazzCost)
	}

	if formData.TeamCompCost != teamCompCost {
		t.Fatalf("found unexpected team competition cost %d, expected %d", formData.TeamCompCost, teamCompCost)
	}

	if formData.TShirtCost != tShirtCost {
		t.Fatalf("found unexpected team competition cost %d, expected %d", formData.TShirtCost, tShirtCost)
	}
}

func TestPopulateMissingItem(t *testing.T) {
	missingItems := []struct {
		Id   string
		Name string
	}{
		{soloJazzItemId, utility.SoloJazzItem},
		{mixAndMatchItemId, utility.MixAndMatchItem},
		{teamCompItemId, utility.TeamCompItem},
		{dancePassItemId, utility.DancePassItem},
		{weekendPassItemId, utility.WeekendPassItem},
		{tShirtItemId, utility.TShirtItem},
	}

	for _, missingItem := range missingItems {
		t.Run(missingItem.Name, func(t *testing.T) {
			logger := logrus.New()
			logger.SetOutput(&test_utility.ErrorWriter{t})
			logger.SetLevel(logrus.TraceLevel)

			catalogObjectsIdx := -1

			incorrectCatalogObjects := []*square.CatalogObject{}
			for _, catalogObject := range catalogObjects {
				if catalogObject.Id != missingItem.Id {
					incorrectCatalogObjects = append(incorrectCatalogObjects, catalogObject)
				}
			}

			client := &mockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					return &mockListCatalogIterator{
						ValueFunc: func() *square.CatalogObject {
							return incorrectCatalogObjects[catalogObjectsIdx]
						},
						ErrorFunc: func() error {
							return nil
						},
						NextFunc: func() bool {
							catalogObjectsIdx++
							return catalogObjectsIdx < len(incorrectCatalogObjects)
						},
					}
				},
				BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
					t.Fatal("Found unexpected call to batch retrieve inventory counts")
					return nil
				},
			}

			_, err := NewService(logger, client).Populate(context.Background())
			if err == nil {
				t.Fatalf("expected error, found none")
			}

			terr, ok := err.(*ErrMissingCatalogItem)
			if !ok {
				t.Fatalf("error was not of type missing catalog item (contains message %v)", err)
			}

			if terr.Name != missingItem.Name {
				t.Fatalf("found missing item %s, when expected missing item %s", terr.Name, missingItem.Name)
			}
		})
	}
}

func TestPopulateNoVariations(t *testing.T) {
	missingItems := []struct {
		Id   string
		Name string
	}{
		{soloJazzItemId, utility.SoloJazzItem},
		{mixAndMatchItemId, utility.MixAndMatchItem},
		{teamCompItemId, utility.TeamCompItem},
		{tShirtItemId, utility.TShirtItem},
	}

	for _, missingItem := range missingItems {
		t.Run(missingItem.Name, func(t *testing.T) {
			logger := logrus.New()
			logger.SetOutput(&test_utility.ErrorWriter{t})
			logger.SetLevel(logrus.TraceLevel)

			catalogObjectsIdx := -1

			incorrectCatalogObjects := []*square.CatalogObject{}
			for _, catalogObject := range catalogObjects {
				if catalogObject.Id != missingItem.Id {
					incorrectCatalogObjects = append(incorrectCatalogObjects, catalogObject)
					continue
				}

				item, ok := catalogObject.CatalogObjectType.(*square.CatalogItem)
				if !ok {
					t.Fatalf("found a top-level catalog object not of type catalog item?")
				}
				incorrectItem := &square.CatalogItem{}
				*incorrectItem = *item
				incorrectItem.Variations = nil

				incorrectObject := &square.CatalogObject{}
				*incorrectObject = *catalogObject
				incorrectObject.CatalogObjectType = incorrectItem
				incorrectCatalogObjects = append(incorrectCatalogObjects, incorrectObject)
			}

			client := &mockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					return &mockListCatalogIterator{
						ValueFunc: func() *square.CatalogObject {
							return incorrectCatalogObjects[catalogObjectsIdx]
						},
						ErrorFunc: func() error {
							return nil
						},
						NextFunc: func() bool {
							catalogObjectsIdx++
							return catalogObjectsIdx < len(incorrectCatalogObjects)
						},
					}
				},
				BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
					t.Fatal("Found unexpected call to batch retrieve inventory counts")
					return nil
				},
			}

			_, err := NewService(logger, client).Populate(context.Background())
			if err == nil {
				t.Fatalf("expected error, found none")
			}

			terr, ok := err.(*ErrUnxpectedVariationCount)
			if !ok {
				t.Fatalf("error was not of type missing catalog item (contains message %v)", err)
			}

			if terr.Name != missingItem.Name {
				t.Fatalf("found item with missing variations %s, when expected item %s", terr.Name, missingItem.Name)
			}

			if terr.Count != 0 {
				t.Fatalf("found incorrect number of variations %d, when none was expected", terr.Count)
			}
		})
	}
}

func TestPopulateMissingVariation(t *testing.T) {
	missingItems := []struct {
		Id            string
		Name          string
		VariationId   string
		VariationName string
	}{
		{dancePassItemId, utility.DancePassItem, dancePassVariationId, utility.DancePassPresaleName},
		{weekendPassItemId, utility.WeekendPassItem, weekendPassTier1VariationId, utility.WeekendPassTier1Name},
		{weekendPassItemId, utility.WeekendPassItem, weekendPassTier2VariationId, utility.WeekendPassTier2Name},
		{weekendPassItemId, utility.WeekendPassItem, weekendPassTier3VariationId, utility.WeekendPassTier3Name},
		{weekendPassItemId, utility.WeekendPassItem, weekendPassTier4VariationId, utility.WeekendPassTier4Name},
		{weekendPassItemId, utility.WeekendPassItem, weekendPassTier5VariationId, utility.WeekendPassTier5Name},
	}

	for _, missingItem := range missingItems {
		t.Run(missingItem.VariationName, func(t *testing.T) {
			logger := logrus.New()
			logger.SetOutput(&test_utility.ErrorWriter{t})
			logger.SetLevel(logrus.TraceLevel)

			catalogObjectsIdx := -1

			incorrectCatalogObjects := []*square.CatalogObject{}
			for _, catalogObject := range catalogObjects {
				if catalogObject.Id != missingItem.Id {
					incorrectCatalogObjects = append(incorrectCatalogObjects, catalogObject)
					continue
				}

				item, ok := catalogObject.CatalogObjectType.(*square.CatalogItem)
				if !ok {
					t.Fatalf("found a top-level catalog object not of type catalog item?")
				}
				incorrectItem := &square.CatalogItem{}
				*incorrectItem = *item
				incorrectItem.Variations = []*square.CatalogObject{}
				for _, variation := range item.Variations {
					if variation.Id != missingItem.VariationId {
						incorrectItem.Variations = append(incorrectItem.Variations, variation)
					}
				}

				incorrectObject := &square.CatalogObject{}
				*incorrectObject = *catalogObject
				incorrectObject.CatalogObjectType = incorrectItem
				incorrectCatalogObjects = append(incorrectCatalogObjects, incorrectObject)
			}

			client := &mockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					return &mockListCatalogIterator{
						ValueFunc: func() *square.CatalogObject {
							return incorrectCatalogObjects[catalogObjectsIdx]
						},
						ErrorFunc: func() error {
							return nil
						},
						NextFunc: func() bool {
							catalogObjectsIdx++
							return catalogObjectsIdx < len(incorrectCatalogObjects)
						},
					}
				},
				BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
					t.Fatal("Found unexpected call to batch retrieve inventory counts")
					return nil
				},
			}

			_, err := NewService(logger, client).Populate(context.Background())
			if err == nil {
				t.Fatalf("expected error, found none")
			}

			terr, ok := err.(*ErrMissingVariation)
			if !ok {
				t.Fatalf("error was not of type missing catalog item (contains message %v)", err)
			}

			if terr.Name != missingItem.Name {
				t.Fatalf("found item with missing variations %s, when expected item %s", terr.Name, missingItem.Name)
			}

			if terr.VariationName != missingItem.VariationName {
				t.Fatalf("found error claiming missing variation name %s, when %s should be missing", terr.VariationName, missingItem.VariationName)
			}
		})
	}
}

func TestPopulateListCatalogError(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	testError := errors.New("some error")

	client := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					t.Fatalf("tried to retrieve a value from an errored iterator")
					return nil
				},
				ErrorFunc: func() error {
					return testError
				},
				NextFunc: func() bool {
					return false
				},
			}
		},
		BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
			t.Fatal("Found unexpected call to batch retrieve inventory counts")
			return nil
		},
	}

	_, err := NewService(logger, client).Populate(context.Background())
	if err == nil {
		t.Fatalf("expected error, found none")
	}

	if errors.Cause(err) != testError {
		t.Fatalf("returned error different than expected (message %v)", err)
	}
}

func TestPopulateBatchCountsError(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	testError := errors.New("some error")
	catalogObjectsIdx := -1

	client := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
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
		},
		BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
			return &mockBatchRetrieveInventoryCountsIterator{
				ValueFunc: func() *square.InventoryCount {
					t.Fatalf("tried to retrieve a value from an errored iterator")
					return nil
				},
				ErrorFunc: func() error {
					return testError
				},
				NextFunc: func() bool {
					return false
				},
			}
		},
	}

	_, err := NewService(logger, client).Populate(context.Background())
	if err == nil {
		t.Fatalf("expected error, found none")
	}

	if errors.Cause(err) != testError {
		t.Fatalf("returned error different than expected (message %v)", err)
	}
}
