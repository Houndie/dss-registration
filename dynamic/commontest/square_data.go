package commontest

import (
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
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
	MixAndMatchID     map[storage.MixAndMatchRole]string
	SoloJazzID        string
	TeamCompetitionID string
	TShirtID          map[storage.TShirtStyle]string

	StudentDiscountID       string
	FullWeekendDiscountID   string
	MixAndMatchDiscountID   string
	FullWeekendDiscountName string
	MixAndMatchDiscountName string
}

func (o *CatalogObjects) SquareData() *common.SquareData {
	return &common.SquareData{
		PurchaseItems: &common.PurchaseItems{
			SoloJazz: &common.PurchaseItem{
				ID:    o.SoloJazzID,
				Price: o.SoloJazzCost,
			},
			MixAndMatch: map[storage.MixAndMatchRole]*common.PurchaseItem{
				storage.MixAndMatchRoleLeader: &common.PurchaseItem{
					ID:    o.MixAndMatchID[storage.MixAndMatchRoleLeader],
					Price: o.MixAndMatchCost,
				},
				storage.MixAndMatchRoleFollower: &common.PurchaseItem{
					ID:    o.MixAndMatchID[storage.MixAndMatchRoleFollower],
					Price: o.MixAndMatchCost,
				},
			},
			TeamCompetition: &common.PurchaseItem{
				ID:    o.TeamCompetitionID,
				Price: o.TeamCompCost,
			},
			TShirt: map[storage.TShirtStyle]*common.PurchaseItem{
				storage.TShirtStyleUnisexS: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisexS],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleUnisexM: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisexM],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleUnisexL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisexL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleUnisexXL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisexXL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleUnisex2XL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisex2XL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleUnisex3XL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleUnisex3XL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleBellaS: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleBellaS],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleBellaM: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleBellaM],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleBellaL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleBellaL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleBellaXL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleBellaXL],
					Price: o.TShirtCost,
				},
				storage.TShirtStyleBella2XL: &common.PurchaseItem{
					ID:    o.TShirtID[storage.TShirtStyleBella2XL],
					Price: o.TShirtCost,
				},
			},
			DanceOnly: &common.PurchaseItem{
				ID:    o.DancePassID,
				Price: o.DancePassCost,
			},
			FullWeekend: map[storage.WeekendPassTier]*common.PurchaseItem{
				storage.Tier1: &common.PurchaseItem{
					ID:    o.WeekendPassID[storage.Tier1],
					Price: o.WeekendPassCost[storage.Tier1],
				},
				storage.Tier2: &common.PurchaseItem{
					ID:    o.WeekendPassID[storage.Tier2],
					Price: o.WeekendPassCost[storage.Tier2],
				},
				storage.Tier3: &common.PurchaseItem{
					ID:    o.WeekendPassID[storage.Tier3],
					Price: o.WeekendPassCost[storage.Tier3],
				},
				storage.Tier4: &common.PurchaseItem{
					ID:    o.WeekendPassID[storage.Tier4],
					Price: o.WeekendPassCost[storage.Tier4],
				},
				storage.Tier5: &common.PurchaseItem{
					ID:    o.WeekendPassID[storage.Tier5],
					Price: o.WeekendPassCost[storage.Tier5],
				},
			},
		},
		Discounts: &common.Discounts{
			StudentDiscount: &common.Discount{
				ID:        o.StudentDiscountID,
				Amount:    common.DollarDiscount(o.StudentDiscountAmount),
				AppliedTo: storage.FullWeekendPurchaseItem,
			},
			CodeDiscounts: map[string][]*common.Discount{
				o.FullWeekendDiscountName: []*common.Discount{
					&common.Discount{
						ID:        o.FullWeekendDiscountID,
						Amount:    common.DollarDiscount(o.FullWeekendDiscountAmount),
						AppliedTo: storage.FullWeekendPurchaseItem,
					},
				},
				o.MixAndMatchDiscountName: []*common.Discount{
					&common.Discount{
						ID:        o.MixAndMatchDiscountID,
						Amount:    common.DollarDiscount(o.MixAndMatchDiscountAmount),
						AppliedTo: storage.MixAndMatchPurchaseItem,
					},
				},
			},
		},
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

		DancePassID: "dance pass id",
		MixAndMatchID: map[storage.MixAndMatchRole]string{
			storage.MixAndMatchRoleLeader:   "mix and match leader id",
			storage.MixAndMatchRoleFollower: "mix and match leader id",
		},
		SoloJazzID:        "solo jazz id",
		TeamCompetitionID: "team competition id",
		TShirtID: map[storage.TShirtStyle]string{
			storage.TShirtStyleUnisexS:   "tshirt unisex s id",
			storage.TShirtStyleUnisexM:   "tshirt unisex m id",
			storage.TShirtStyleUnisexL:   "tshirt unisex l id",
			storage.TShirtStyleUnisexXL:  "tshirt unisex xl id",
			storage.TShirtStyleUnisex2XL: "tshirt unisex 2xl id",
			storage.TShirtStyleUnisex3XL: "tshirt unisex 3xl id",
			storage.TShirtStyleBellaS:    "tshirt bella s id",
			storage.TShirtStyleBellaM:    "tshirt bella m id",
			storage.TShirtStyleBellaL:    "tshirt bella l id",
			storage.TShirtStyleBellaXL:   "tshirt bella xl id",
			storage.TShirtStyleBella2XL:  "tshirt bella 2xl id",
		},

		StudentDiscountID:       "student discount id",
		FullWeekendDiscountID:   "full weekend discount id",
		MixAndMatchDiscountID:   "mix and match discount id",
		FullWeekendDiscountName: "full weekend discount name",
		MixAndMatchDiscountName: "mix and match discount name",
	}
}
