package postgres

import "github.com/Houndie/dss-registration/dynamic/storage"

const (
	discountBundleTable   = "discount_bundles"
	discountBundleIDCol   = "id"
	discountBundleCodeCol = "code"

	discountTable        = "discounts"
	discountFkCol        = "discount_bundle_id"
	discountNameCol      = "name"
	discountAppliedToCol = "applied_to"
)

type discountConstsTypes struct {
	BundleTable   string
	BundleIDCol   string
	BundleCodeCol string

	SDTable        string
	SDFkCol        string
	SDNameCol      string
	SDAppliedToCol string
}

var discountConsts = &discountConstsTypes{
	BundleTable:   "discount_bundles",
	BundleIDCol:   "id",
	BundleCodeCol: "code",

	SDTable:        "discounts",
	SDFkCol:        "discount_bundle_id",
	SDNameCol:      "name",
	SDAppliedToCol: "applied_to",
}

var appliedToToEnum = map[storage.PurchaseItem]string{
	storage.FullWeekendPurchaseItem:     "Full Weekend",
	storage.DanceOnlyPurchaseItem:       "Dance Only",
	storage.MixAndMatchPurchaseItem:     "Mix And Match",
	storage.SoloJazzPurchaseItem:        "Solo Jazz",
	storage.TeamCompetitionPurchaseItem: "Team Competition",
	storage.TShirtPurchaseItem:          "TShirt",
}

var appliedToFromEnum = map[string]storage.PurchaseItem{
	"Full Weekend":     storage.FullWeekendPurchaseItem,
	"Dance Only":       storage.DanceOnlyPurchaseItem,
	"Mix And Match":    storage.MixAndMatchPurchaseItem,
	"Solo Jazz":        storage.SoloJazzPurchaseItem,
	"Team Competition": storage.TeamCompetitionPurchaseItem,
	"TShirt":           storage.TShirtPurchaseItem,
}
