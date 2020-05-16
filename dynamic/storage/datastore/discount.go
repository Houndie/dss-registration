package datastore

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

const (
	discountKind = "Discount"

	fullWeekendDiscount     = "full weekend"
	danceOnlyDiscount       = "dance only"
	mixAndMatchDiscount     = "mix and match"
	soloJazzDiscount        = "solo jazz"
	teamCompetitionDiscount = "team competition"
	tshirtDiscount          = "tshirt"
)

type singleDiscount struct {
	Name      string
	AppliedTo string
}

type discountEntity struct {
	Code      string
	Discounts []singleDiscount `datastore:",noindex"`
}

func toDiscountEntity(d *storage.Discount) (*discountEntity, error) {
	singleDiscounts := make([]singleDiscount, len(d.Discounts))
	for i, sd := range d.Discounts {
		var appliedTo string
		switch sd.AppliedTo {
		case storage.FullWeekendPurchaseItem:
			appliedTo = fullWeekendDiscount
		case storage.DanceOnlyPurchaseItem:
			appliedTo = danceOnlyDiscount
		case storage.MixAndMatchPurchaseItem:
			appliedTo = mixAndMatchDiscount
		case storage.SoloJazzPurchaseItem:
			appliedTo = soloJazzDiscount
		case storage.TeamCompetitionPurchaseItem:
			appliedTo = teamCompetitionDiscount
		case storage.TShirtPurchaseItem:
			appliedTo = tshirtDiscount
		default:
			return nil, fmt.Errorf("unknown purchase item found: %v", sd.AppliedTo)
		}

		singleDiscounts[i] = singleDiscount{
			Name:      sd.Name,
			AppliedTo: appliedTo,
		}
	}

	return &discountEntity{
		Code:      d.Code,
		Discounts: singleDiscounts,
	}, nil
}

func fromDiscountEntity(key *datastore.Key, de *discountEntity) (*storage.Discount, error) {
	singleDiscounts := make([]*storage.SingleDiscount, len(de.Discounts))
	for i, sd := range de.Discounts {
		var appliedTo storage.PurchaseItem
		switch sd.AppliedTo {
		case fullWeekendDiscount:
			appliedTo = storage.FullWeekendPurchaseItem
		case danceOnlyDiscount:
			appliedTo = storage.DanceOnlyPurchaseItem
		case mixAndMatchDiscount:
			appliedTo = storage.MixAndMatchPurchaseItem
		case soloJazzDiscount:
			appliedTo = storage.SoloJazzPurchaseItem
		case teamCompetitionDiscount:
			appliedTo = storage.TeamCompetitionPurchaseItem
		case tshirtDiscount:
			appliedTo = storage.TShirtPurchaseItem
		default:
			return nil, fmt.Errorf("unknown purchase item found: %v", sd.AppliedTo)
		}

		singleDiscounts[i] = &storage.SingleDiscount{
			Name:      sd.Name,
			AppliedTo: appliedTo,
		}
	}

	return &storage.Discount{
		ID:        key.Encode(),
		Code:      de.Code,
		Discounts: singleDiscounts,
	}, nil
}
