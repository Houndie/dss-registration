package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/adddiscount"
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
		case adddiscount.FullWeekendDiscountTarget:
			appliedTo = fullWeekendDiscount
		case adddiscount.DanceOnlyDiscountTarget:
			appliedTo = danceOnlyDiscount
		case adddiscount.MixAndMatchDiscountTarget:
			appliedTo = mixAndMatchDiscount
		case adddiscount.SoloJazzDiscountTarget:
			appliedTo = soloJazzDiscount
		case adddiscount.TeamCompetitionDiscountTarget:
			appliedTo = teamCompetitionDiscount
		case adddiscount.TShirtDiscountTarget:
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
