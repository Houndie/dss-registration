package add

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

	mixAndMatchItemId      = "mix and match id"
	mixAndMatchVariationId = "mix and match variation id"

	teamCompItemId      = "team competition id"
	teamCompVariationId = "team competition variation id"

	dancePassItemId      = "dance pass id"
	dancePassVariationId = "dance pass variation id"

	weekendPassItemId           = "weekend pass id"
	weekendPassTier1VariationId = "weekend pass tier 1 variation id"
	weekendPassTier2VariationId = "weekend pass tier 2 variation id"
	weekendPassTier3VariationId = "weekend pass tier 3 variation id"
	weekendPassTier4VariationId = "weekend pass tier 4 variation id"
	weekendPassTier5VariationId = "weekend pass tier 5 variation id"

	tShirtItemId      = "tshirt id"
	tShirtVariationId = "tshirt variation id"
)

var (
	catalogObjects  []*square.CatalogObject
	inventoryCounts []*square.InventoryCount
	locations       []*square.Location
)

func TestMain(m *testing.M) {
	locationCreatedAt := time.Unix(200, 0)
	locationStartTime := time.Unix(201, 0)
	locationEndTime := time.Unix(202, 0)
	locations = []*square.Location{
		&square.Location{
			Id:           "some location id",
			Name:         "THE INTERNET",
			Address:      nil,
			Timezone:     "Eastern", //TODO Enum the timezones?
			Capabilities: []square.LocationCapability{square.LocationCapabilityCreditCardProcessing},
			Status:       square.LocationStatusActive,
			CreatedAt:    &locationCreatedAt,
			MerchantId:   "some merchant id",
			Country:      "USA", //TODO enum this?
			LanguageCode: "EN",  //TODO enum this?
			Currency:     "USD", //TODO enum this
			PhoneNumber:  "1234567890",
			BusinessName: "Dayton Swing Smackdown",
			Type:         square.LocationTypeMobile,
			WebsiteUrl:   "https://www.daytonswingsmackdown.com",
			BusinessHours: &square.BusinessHours{
				Periods: []*square.BusinessHoursPeriod{
					&square.BusinessHoursPeriod{
						DayOfWeek:      "MONDAY",
						StartLocalTime: &locationStartTime,
						EndLocalTime:   &locationEndTime,
					},
				},
			},
			BusinessEmail:     "info@daytonswingsmackdown.com",
			Description:       "we do the swing dance",
			TwitterUsername:   "@daytonswingsmackdown",
			InstagramUsername: "@daytonswingsmackdown",
			FacebookUrl:       "https://facebook.com/daytonswingsmackdown",
			Coordinates: &square.Coordinates{
				Latitude:  7,
				Longitude: 12.0938498734509873049872309872938092389750987230987209834087239823897598234987,
			},
		},
	}
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
								Amount:   500,
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
								Amount:   500,
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
								Amount:   5500,
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
								Amount:   4500,
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
								Amount:   8500,
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
								Amount:   1500,
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
			LocationId:        locations[0].Id,
			Quantity:          "0",
			CalculatedAt:      &weekendPassTier1CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier2VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locations[0].Id,
			Quantity:          "0",
			CalculatedAt:      &weekendPassTier2CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier3VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locations[0].Id,
			Quantity:          "8",
			CalculatedAt:      &weekendPassTier3CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier4VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locations[0].Id,
			Quantity:          "9",
			CalculatedAt:      &weekendPassTier4CalculatedAt,
		},
		&square.InventoryCount{
			CatalogObjectId:   weekendPassTier5VariationId,
			CatalogObjectType: square.CatalogObjectTypeItemVariation,
			State:             square.InventoryStateInStock,
			LocationId:        locations[0].Id,
			Quantity:          "10",
			CalculatedAt:      &weekendPassTier5CalculatedAt,
		},
	}
	os.Exit(m.Run())
}
