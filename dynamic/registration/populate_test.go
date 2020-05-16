package registration

import (
	"context"
	"fmt"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/sirupsen/logrus"
)

type catalogObjects struct {
	soloJazzCost    int
	mixAndMatchCost int
	teamCompCost    int
	dancePassCost   int
	tShirtCost      int
	weekendPassCost map[storage.WeekendPassTier]int

	studentDiscountAmount     int
	fullWeekendDiscountAmount int
	mixAndMatchDiscountAmount int

	weekendPassID     map[storage.WeekendPassTier]string
	dancePassID       string
	mixAndMatchID     string
	soloJazzID        string
	teamCompetitionID string
	tShirtID          string

	studentDiscountID       string
	fullWeekendDiscountID   string
	mixAndMatchDiscountID   string
	fullWeekendDiscountName string
	mixAndMatchDiscountName string
}

func (o *catalogObjects) catalog() []*square.CatalogObject {
	return []*square.CatalogObject{
		soloJazzItem(o.soloJazzID, o.soloJazzCost),
		mixAndMatchItem(o.mixAndMatchID, o.mixAndMatchCost),
		teamCompItem(o.teamCompetitionID, o.teamCompCost),
		tShirtItem(o.tShirtID, o.tShirtCost),
		dancePassItem(o.dancePassID, o.dancePassCost),
		weekendPassItem(o.weekendPassID, o.weekendPassCost),
		discountItem(utility.StudentDiscountItem, o.studentDiscountAmount, o.studentDiscountID),
		discountItem(o.fullWeekendDiscountName, o.fullWeekendDiscountAmount, o.fullWeekendDiscountID),
		discountItem(o.mixAndMatchDiscountName, o.mixAndMatchDiscountAmount, o.mixAndMatchDiscountID),
	}
}

func commonCatalogObjects() *catalogObjects {
	return &catalogObjects{
		soloJazzCost:    1,
		mixAndMatchCost: 2,
		teamCompCost:    3,
		dancePassCost:   4,
		tShirtCost:      5,
		weekendPassCost: map[storage.WeekendPassTier]int{
			storage.Tier1: 6,
			storage.Tier2: 7,
			storage.Tier3: 8,
			storage.Tier4: 9,
			storage.Tier5: 10,
		},

		studentDiscountAmount:     11,
		fullWeekendDiscountAmount: 15,
		mixAndMatchDiscountAmount: 1,

		weekendPassID: map[storage.WeekendPassTier]string{
			storage.Tier1: "weekend pass tier 1 variation id",
			storage.Tier2: "weekend pass tier 2 variation id",
			storage.Tier3: "weekend pass tier 3 variation id",
			storage.Tier4: "weekend pass tier 4 variation id",
			storage.Tier5: "weekend pass tier 5 variation id",
		},
		dancePassID:       "dance pass id",
		mixAndMatchID:     "mix and match id",
		soloJazzID:        "solo jazz id",
		teamCompetitionID: "team competition id",
		tShirtID:          "tshirt id",

		studentDiscountID:       "student discount id",
		fullWeekendDiscountID:   "full weekend discount id",
		mixAndMatchDiscountID:   "mix and match discount id",
		fullWeekendDiscountName: "full weekend discount name",
		mixAndMatchDiscountName: "mix and match discount name",
	}
}

func discountItem(name string, amt int, id string) *square.CatalogObject {
	return &square.CatalogObject{
		Id: id,
		CatalogObjectType: &square.CatalogDiscount{
			Name: name,
			DiscountType: &square.CatalogDiscountFixedAmount{
				AmountMoney: &square.Money{
					Amount: amt,
				},
			},
		},
	}
}

func soloJazzItem(id string, cost int) *square.CatalogObject {
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name: utility.SoloJazzItem,
			Variations: []*square.CatalogObject{
				&square.CatalogObject{
					Id: id,
					CatalogObjectType: &square.CatalogItemVariation{
						PriceMoney: &square.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func mixAndMatchItem(id string, cost int) *square.CatalogObject {
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name: utility.MixAndMatchItem,
			Variations: []*square.CatalogObject{
				&square.CatalogObject{
					Id: id,
					CatalogObjectType: &square.CatalogItemVariation{
						PriceMoney: &square.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func teamCompItem(id string, cost int) *square.CatalogObject {
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name: utility.TeamCompItem,
			Variations: []*square.CatalogObject{
				&square.CatalogObject{
					Id: id,
					CatalogObjectType: &square.CatalogItemVariation{
						PriceMoney: &square.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func tShirtItem(id string, cost int) *square.CatalogObject {
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name: utility.TShirtItem,
			Variations: []*square.CatalogObject{
				&square.CatalogObject{
					Id: id,
					CatalogObjectType: &square.CatalogItemVariation{
						PriceMoney: &square.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func dancePassItem(id string, cost int) *square.CatalogObject {
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name: utility.DancePassItem,
			Variations: []*square.CatalogObject{
				&square.CatalogObject{
					Id: id,
					CatalogObjectType: &square.CatalogItemVariation{
						Name: utility.DancePassPresaleName,
						PriceMoney: &square.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func weekendPassItem(weekendPassIds map[storage.WeekendPassTier]string, weekendPassCosts map[storage.WeekendPassTier]int) *square.CatalogObject {
	variations := make([]*square.CatalogObject, len(weekendPassIds))
	idx := 0
	for tier, id := range weekendPassIds {
		variations[idx] = &square.CatalogObject{
			Id: id,
			CatalogObjectType: &square.CatalogItemVariation{
				Name: utility.WeekendPassName[tier],
				PriceMoney: &square.Money{
					Amount: weekendPassCosts[tier],
				},
			},
		}
		idx++
	}
	return &square.CatalogObject{
		CatalogObjectType: &square.CatalogItem{
			Name:       utility.WeekendPassItem,
			Variations: variations,
		},
	}
}

func TestPopulate(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{T: t})
	logger.AddHook(&test_utility.ErrorHook{T: t})
	logger.SetLevel(logrus.TraceLevel)

	co := commonCatalogObjects()

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

			var expectedCost = co.weekendPassCost[expectTier]

			inventoryCounts := make([]*square.InventoryCount, len(counts))
			idx := 0
			for tier, count := range counts {
				inventoryCounts[idx] = &square.InventoryCount{
					CatalogObjectId: co.weekendPassID[tier],
					Quantity:        count,
				}
				idx++
			}

			client := &mockSquareClient{
				ListCatalogFunc:                  listCatalogFuncFromSlice(co.catalog()),
				BatchRetrieveInventoryCountsFunc: inventoryCountsFromSliceCheck(t, co.weekendPassID, inventoryCounts),
			}

			formData, err := NewService(true, logger, client, &mockAuthorizer{}, &mockStore{}, &mockMailClient{}).Populate(context.Background())
			if err != nil {
				t.Fatalf("error populating form data: %v", err)
			}

			if formData.WeekendPassCost != expectedCost {
				t.Errorf("found unexpected weekend pass cost %d, expected %d", formData.WeekendPassCost, expectedCost)
			}

			if formData.WeekendPassTier != expectTier {
				t.Errorf("found unexpected weekend pass tier %d, expected %d", formData.WeekendPassTier, expectTier)
			}

			if formData.DancePassCost != co.dancePassCost {
				t.Errorf("found unexpected dance pass cost %d, expected %d", formData.DancePassCost, co.dancePassCost)
			}

			if formData.MixAndMatchCost != co.mixAndMatchCost {
				t.Errorf("found unexpected mix and match cost %d, expected %d", formData.MixAndMatchCost, co.mixAndMatchCost)
			}

			if formData.SoloJazzCost != co.soloJazzCost {
				t.Errorf("found unexpected solo jazz cost %d, expected %d", formData.SoloJazzCost, co.soloJazzCost)
			}

			if formData.TeamCompCost != co.teamCompCost {
				t.Errorf("found unexpected team competition cost %d, expected %d", formData.TeamCompCost, co.teamCompCost)
			}

			if formData.TShirtCost != co.tShirtCost {
				t.Errorf("found unexpected team competition cost %d, expected %d", formData.TShirtCost, co.tShirtCost)
			}

			dd, ok := formData.StudentDiscount.(DollarDiscount)
			if !ok {
				t.Fatalf("student disocunt is not of dollar discount type")
			}

			if int(dd) != co.studentDiscountAmount {
				t.Errorf("unexpected student discount amount %d, expected %d", int(dd), co.studentDiscountAmount)
			}
		})
	}

}
