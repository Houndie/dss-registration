package commontest

import (
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/Houndie/square-go/objects"
)

type CatalogObjects struct {
	SoloJazzCost    int
	MixAndMatchCost int
	TeamCompCost    int
	DancePassCost   int
	TShirtCost      int
	WeekendPassCost map[storage.WeekendPassTier]int

	StudentDiscountAmount     int
	FullWeekendDiscountAmount int
	MixAndMatchDiscountAmount int

	WeekendPassID     map[storage.WeekendPassTier]string
	DancePassID       string
	MixAndMatchID     string
	SoloJazzID        string
	TeamCompetitionID string
	TShirtID          string

	StudentDiscountID       string
	FullWeekendDiscountID   string
	MixAndMatchDiscountID   string
	FullWeekendDiscountName string
	MixAndMatchDiscountName string
}

func (o *CatalogObjects) Catalog() []*objects.CatalogObject {
	return []*objects.CatalogObject{
		soloJazzItem(o.SoloJazzID, o.SoloJazzCost),
		mixAndMatchItem(o.MixAndMatchID, o.MixAndMatchCost),
		teamCompItem(o.TeamCompetitionID, o.TeamCompCost),
		tShirtItem(o.TShirtID, o.TShirtCost),
		dancePassItem(o.DancePassID, o.DancePassCost),
		weekendPassItem(o.WeekendPassID, o.WeekendPassCost),
		discountItem(utility.StudentDiscountItem, o.StudentDiscountAmount, o.StudentDiscountID),
		discountItem(o.FullWeekendDiscountName, o.FullWeekendDiscountAmount, o.FullWeekendDiscountID),
		discountItem(o.MixAndMatchDiscountName, o.MixAndMatchDiscountAmount, o.MixAndMatchDiscountID),
	}
}

func CommonCatalogObjects() *CatalogObjects {
	return &CatalogObjects{
		SoloJazzCost:    1,
		MixAndMatchCost: 2,
		TeamCompCost:    3,
		DancePassCost:   4,
		TShirtCost:      5,
		WeekendPassCost: map[storage.WeekendPassTier]int{
			storage.Tier1: 6,
			storage.Tier2: 7,
			storage.Tier3: 8,
			storage.Tier4: 9,
			storage.Tier5: 10,
		},

		StudentDiscountAmount:     11,
		FullWeekendDiscountAmount: 15,
		MixAndMatchDiscountAmount: 1,

		WeekendPassID: map[storage.WeekendPassTier]string{
			storage.Tier1: "weekend pass tier 1 variation id",
			storage.Tier2: "weekend pass tier 2 variation id",
			storage.Tier3: "weekend pass tier 3 variation id",
			storage.Tier4: "weekend pass tier 4 variation id",
			storage.Tier5: "weekend pass tier 5 variation id",
		},
		DancePassID:       "dance pass id",
		MixAndMatchID:     "mix and match id",
		SoloJazzID:        "solo jazz id",
		TeamCompetitionID: "team competition id",
		TShirtID:          "tshirt id",

		StudentDiscountID:       "student discount id",
		FullWeekendDiscountID:   "full weekend discount id",
		MixAndMatchDiscountID:   "mix and match discount id",
		FullWeekendDiscountName: "full weekend discount name",
		MixAndMatchDiscountName: "mix and match discount name",
	}
}

func discountItem(name string, amt int, id string) *objects.CatalogObject {
	return &objects.CatalogObject{
		ID: id,
		Type: &objects.CatalogDiscount{
			Name: name,
			DiscountType: &objects.CatalogDiscountFixedAmount{
				AmountMoney: &objects.Money{
					Amount: amt,
				},
			},
		},
	}
}

func soloJazzItem(id string, cost int) *objects.CatalogObject {
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name: utility.SoloJazzItem,
			Variations: []*objects.CatalogObject{
				&objects.CatalogObject{
					ID: id,
					Type: &objects.CatalogItemVariation{
						PriceMoney: &objects.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func mixAndMatchItem(id string, cost int) *objects.CatalogObject {
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name: utility.MixAndMatchItem,
			Variations: []*objects.CatalogObject{
				&objects.CatalogObject{
					ID: id,
					Type: &objects.CatalogItemVariation{
						PriceMoney: &objects.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func teamCompItem(id string, cost int) *objects.CatalogObject {
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name: utility.TeamCompItem,
			Variations: []*objects.CatalogObject{
				&objects.CatalogObject{
					ID: id,
					Type: &objects.CatalogItemVariation{
						PriceMoney: &objects.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func tShirtItem(id string, cost int) *objects.CatalogObject {
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name: utility.TShirtItem,
			Variations: []*objects.CatalogObject{
				&objects.CatalogObject{
					ID: id,
					Type: &objects.CatalogItemVariation{
						PriceMoney: &objects.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func dancePassItem(id string, cost int) *objects.CatalogObject {
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name: utility.DancePassItem,
			Variations: []*objects.CatalogObject{
				&objects.CatalogObject{
					ID: id,
					Type: &objects.CatalogItemVariation{
						Name: utility.DancePassPresaleName,
						PriceMoney: &objects.Money{
							Amount: cost,
						},
					},
				},
			},
		},
	}
}

func weekendPassItem(weekendPassIDs map[storage.WeekendPassTier]string, weekendPassCosts map[storage.WeekendPassTier]int) *objects.CatalogObject {
	variations := make([]*objects.CatalogObject, len(weekendPassIDs))
	idx := 0
	for tier, id := range weekendPassIDs {
		variations[idx] = &objects.CatalogObject{
			ID: id,
			Type: &objects.CatalogItemVariation{
				Name: utility.WeekendPassName[tier],
				PriceMoney: &objects.Money{
					Amount: weekendPassCosts[tier],
				},
			},
		}
		idx++
	}
	return &objects.CatalogObject{
		Type: &objects.CatalogItem{
			Name:       utility.WeekendPassItem,
			Variations: variations,
		},
	}
}
