package registration

import (
	"context"
	"fmt"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/objects"
	"github.com/sirupsen/logrus"
)

func TestPopulate(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{T: t})
	logger.AddHook(&test_utility.ErrorHook{T: t})
	logger.SetLevel(logrus.TraceLevel)

	co := commontest.CommonCatalogObjects()

	expectTiers := []storage.WeekendPassTier{storage.Tier1, storage.Tier2, storage.Tier3, storage.Tier4, storage.Tier5}
	for _, expectTier := range expectTiers {
		t.Run(fmt.Sprintf("expect_tier_%d", expectTier), func(t *testing.T) {
			var counts map[storage.WeekendPassTier]string
			switch expectTier {
			case storage.Tier1:
				counts = map[storage.WeekendPassTier]string{storage.Tier1: "25", storage.Tier2: "25", storage.Tier3: "25", storage.Tier4: "25", storage.Tier5: "25"}
			case storage.Tier2:
				counts = map[storage.WeekendPassTier]string{storage.Tier1: "0", storage.Tier2: "25", storage.Tier3: "25", storage.Tier4: "25", storage.Tier5: "25"}
			case storage.Tier3:
				counts = map[storage.WeekendPassTier]string{storage.Tier1: "0", storage.Tier2: "0", storage.Tier3: "25", storage.Tier4: "25", storage.Tier5: "25"}
			case storage.Tier4:
				counts = map[storage.WeekendPassTier]string{storage.Tier1: "0", storage.Tier2: "0", storage.Tier3: "0", storage.Tier4: "25", storage.Tier5: "25"}
			case storage.Tier5:
				counts = map[storage.WeekendPassTier]string{storage.Tier1: "0", storage.Tier2: "0", storage.Tier3: "0", storage.Tier4: "0", storage.Tier5: "25"}
			}

			inventoryCounts := make([]*objects.InventoryCount, len(counts))
			idx := 0
			for tier, count := range counts {
				inventoryCounts[idx] = &objects.InventoryCount{
					CatalogObjectID: co.WeekendPassID[tier],
					Quantity:        count,
					State:           objects.InventoryStateInStock,
				}
				idx++
			}

			client := &square.Client{
				Inventory: &commontest.MockSquareInventoryClient{
					BatchRetrieveCountsFunc: commontest.InventoryCountsFromSliceCheck(t, co.WeekendPassID, inventoryCounts),
				},
			}

			tier, err := NewService(true, false, logger, client, commontest.CommonCatalogObjects().SquareData(), &commontest.MockAuthorizer{}, &commontest.MockStore{}, &commontest.MockMailClient{}, testPermissionConfig).Populate(context.Background())
			if err != nil {
				t.Fatalf("error populating form data: %v", err)
			}

			if tier != expectTier {
				t.Errorf("found unexpected weekend pass tier %d, expected %d", tier, expectTier)
			}

		})
	}

}
