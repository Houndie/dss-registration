package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

func (d *Datastore) AddDiscount(ctx context.Context, discount *storage.Discount) error {
	storeDiscount := discountEntity{
		Code:      discount.Code,
		Discounts: make([]singleDiscount, len(discount.Discounts)),
	}
	for i, sd := range discount.Discounts {
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
			return errors.New("Unknown discount target found")
		}

		storeDiscount.Discounts[i] = singleDiscount{
			Name:      sd.Name,
			AppliedTo: appliedTo,
		}
	}

	key := datastore.IncompleteKey(discountKind, nil)
	_, err := d.client.Put(ctx, key, &storeDiscount)
	if err != nil {
		return errors.Wrap(err, "error inserting discount into datastore")
	}
	return nil
}
