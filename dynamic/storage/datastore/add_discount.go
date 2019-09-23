package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/adddiscount"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/pkg/errors"
)

func (d *Datastore) AddDiscount(ctx context.Context, discount *adddiscount.Discount) error {
	storeDiscount := discountEntity{
		Code:      discount.Code,
		Discounts: make([]singleDiscount, len(discount.Discounts)),
	}
	for i, sd := range discount.Discounts {
		var appliedTo string
		switch sd.AppliedTo {
		case common.FullWeekendPurchaseItem:
			appliedTo = fullWeekendDiscount
		case common.DanceOnlyPurchaseItem:
			appliedTo = danceOnlyDiscount
		case common.MixAndMatchPurchaseItem:
			appliedTo = mixAndMatchDiscount
		case common.SoloJazzPurchaseItem:
			appliedTo = soloJazzDiscount
		case common.TeamCompetitionPurchaseItem:
			appliedTo = teamCompetitionDiscount
		case common.TShirtPurchaseItem:
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
