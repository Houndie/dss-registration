package commontest

import (
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
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

func (o *CatalogObjects) Catalog() []*square.CatalogObject {
	return []*square.CatalogObject{
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
