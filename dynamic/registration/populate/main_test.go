package populate

import (
	"os"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
)

const (
	soloJazzItemId      = "solo jazz id"
	soloJazzVariationId = "solo jazz variation id"
	soloJazzCost        = 1

	mixAndMatchItemId      = "mix and match id"
	mixAndMatchVariationId = "mix and match variation id"
	mixAndMatchCost        = 2

	teamCompItemId      = "team competition id"
	teamCompVariationId = "team competition variation id"
	teamCompCost        = 3

	dancePassItemId      = "dance pass id"
	dancePassVariationId = "dance pass variation id"
	dancePassCost        = 4

	weekendPassItemId           = "weekend pass id"
	weekendPassTier1VariationId = "weekend pass tier 1 variation id"
	weekendPassTier2VariationId = "weekend pass tier 2 variation id"
	weekendPassTier3VariationId = "weekend pass tier 3 variation id"
	weekendPassTier4VariationId = "weekend pass tier 4 variation id"
	weekendPassTier5VariationId = "weekend pass tier 5 variation id"
	weekendPassCost             = 5

	tShirtItemId      = "tshirt id"
	tShirtVariationId = "tshirt variation id"
	tShirtCost        = 6
)

var catalogObjects []*square.CatalogObject
var inventoryCounts []*square.InventoryCount

func TestMain(m *testing.M) {
	soloJazzUpdatedAt := time.Unix(0, 0)
	soloJazzVariationUpdatedAt := time.Unix(1, 0)

	mixAndMatchUpdatedAt := time.Unix(2, 0)
	mixAndMatchVariationUpdatedAt := time.Unix(3, 0)

	teamCompUpdatedAt := time.Unix(4, 0)
	teamCompVariationUpdatedAt := time.Unix(5, 0)

	dancePassUpdatedAt := time.Unix(6, 0)
	dancePassBadVariationUpdatedAt := time.Unix(7, 0)
	dancePassGoodVariationUpdatedAt := time.Unix(8, 0)

	weekendPassUpdatedAt := time.Unix(9, 0)
	weekendPassBadVariationUpdatedAt := time.Unix(10, 0)
	weekendPassTier1VariationUpdatedAt := time.Unix(11, 0)
	weekendPassTier2VariationUpdatedAt := time.Unix(12, 0)
	weekendPassTier3VariationUpdatedAt := time.Unix(13, 0)
	weekendPassTier4VariationUpdatedAt := time.Unix(14, 0)
	weekendPassTier5VariationUpdatedAt := time.Unix(15, 0)

	tShirtUpdatedAt := time.Unix(16, 0)
	tShirtVariationUpdatedAt := time.Unix(17, 0)

	catalogObjects = []*square.CatalogObject{
		&square.CatalogObject{
			Id:                    soloJazzItemId,
			UpdatedAt:             &soloJazzUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "solo jazz image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.SoloJazzItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    soloJazzVariationId,
						UpdatedAt:             &soloJazzVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "solo jazz variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      soloJazzItemId,
							Name:        "Regular",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   soloJazzCost,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
		&square.CatalogObject{
			Id:                    mixAndMatchItemId,
			UpdatedAt:             &mixAndMatchUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "mix and match image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.MixAndMatchItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    mixAndMatchVariationId,
						UpdatedAt:             &mixAndMatchVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "mix and match variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      mixAndMatchItemId,
							Name:        "Regular",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   mixAndMatchCost,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
		&square.CatalogObject{
			Id:                    teamCompItemId,
			UpdatedAt:             &teamCompUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "team competition image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.TeamCompItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    teamCompVariationId,
						UpdatedAt:             &teamCompVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "team competition variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      teamCompItemId,
							Name:        "Regular",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   teamCompCost,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
		&square.CatalogObject{
			Id:                    dancePassItemId,
			UpdatedAt:             &dancePassUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "dance pass image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.DancePassItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    "dance pass bad variation id",
						UpdatedAt:             &dancePassBadVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "dance pass bad variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      dancePassItemId,
							Name:        "Door",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   5000,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    dancePassVariationId,
						UpdatedAt:             &dancePassGoodVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "dance pass good variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      dancePassItemId,
							Name:        utility.DancePassPresaleName,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   dancePassCost,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
		&square.CatalogObject{
			Id:                    weekendPassItemId,
			UpdatedAt:             &weekendPassUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "weekend pass image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.WeekendPassItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    "weekend pass bad variation id",
						UpdatedAt:             &weekendPassBadVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass bad variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        "Door",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   11000,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    weekendPassTier1VariationId,
						UpdatedAt:             &weekendPassTier1VariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass tier 1 variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        utility.WeekendPassTier1Name,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   6500,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    weekendPassTier2VariationId,
						UpdatedAt:             &weekendPassTier2VariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass tier 2 variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        utility.WeekendPassTier2Name,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   7500,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    weekendPassTier3VariationId,
						UpdatedAt:             &weekendPassTier3VariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass tier 3 variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        utility.WeekendPassTier3Name,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   weekendPassCost,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    weekendPassTier4VariationId,
						UpdatedAt:             &weekendPassTier4VariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass tier 4 variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        utility.WeekendPassTier4Name,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   9500,
								Currency: "USD",
							},
						},
					},
					&square.CatalogObject{
						Id:                    weekendPassTier5VariationId,
						UpdatedAt:             &weekendPassTier5VariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "weekend pass tier 5 variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      weekendPassItemId,
							Name:        utility.WeekendPassTier5Name,
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   10500,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
		&square.CatalogObject{
			Id:                    tShirtItemId,
			UpdatedAt:             &tShirtUpdatedAt,
			Version:               1,
			PresentAtAllLocations: true,
			ImageId:               "tshirt image id",
			CatalogObjectType: &square.CatalogItem{
				Name:                    utility.TShirtItem,
				AvailableOnline:         true,
				AvailableElectronically: true,
				Variations: []*square.CatalogObject{
					&square.CatalogObject{
						Id:                    tShirtVariationId,
						UpdatedAt:             &tShirtVariationUpdatedAt,
						Version:               1,
						PresentAtAllLocations: true,
						ImageId:               "teshirt variation image id",
						CatalogObjectType: &square.CatalogItemVariation{
							ItemId:      tShirtItemId,
							Name:        "Regular",
							PricingType: square.CatalogPricingTypeFixed,
							PriceMoney: &square.Money{
								Amount:   tShirtCost,
								Currency: "USD",
							},
						},
					},
				},
				ProductType:        square.CatalogItemProductTypeRegular,
				SkipModifierScreen: true,
			},
		},
	}

	locationId := "online location id"

	weekendPassTier1CalculatedAt := time.Unix(105, 0)
	weekendPassTier2CalculatedAt := time.Unix(106, 0)
	weekendPassTier3CalculatedAt := time.Unix(107, 0)
	weekendPassTier4CalculatedAt := time.Unix(108, 0)
	weekendPassTier5CalculatedAt := time.Unix(109, 0)

	inventoryCounts = []*square.InventoryCount{
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier1VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locationId,
			Quantity:          "0",
			CalculatedAt:      &weekendPassTier1CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier2VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locationId,
			Quantity:          "0",
			CalculatedAt:      &weekendPassTier2CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier3VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locationId,
			Quantity:          "8",
			CalculatedAt:      &weekendPassTier3CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier4VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locationId,
			Quantity:          "9",
			CalculatedAt:      &weekendPassTier4CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier5VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locationId,
			Quantity:          "10",
			CalculatedAt:      &weekendPassTier5CalculatedAt,
		},
	}
	os.Exit(m.Run())
}
