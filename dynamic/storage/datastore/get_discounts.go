package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/pkg/errors"
)

func (d *Datastore) GetDiscounts(ctx context.Context, codes []string) ([]string, []*common.StoreDiscount, error) {
	result := []*common.StoreDiscount{}
	keys := []string{}
	for _, code := range codes {
		q := datastore.NewQuery(discountKind).Filter("Code =", code).Limit(1)
		discounts := []discountEntity{}
		key, err := d.client.GetAll(ctx, q, &discounts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "error fetching discounts with code %s from datastore", code)
		}

		if len(discounts) == 0 {
			return nil, nil, add.ErrDiscountDoesNotExist{
				Code: code,
			}
		}

		if len(discounts) > 1 {
			return nil, nil, fmt.Errorf("somehow discovered %d discounts with code %s when only one was expected", len(discounts), code)
		}

		for _, sd := range discounts[0].Discounts {
			var appliedTo common.PurchaseItem
			switch sd.AppliedTo {
			case fullWeekendDiscount:
				appliedTo = common.FullWeekendPurchaseItem
			case danceOnlyDiscount:
				appliedTo = common.DanceOnlyPurchaseItem
			case mixAndMatchDiscount:
				appliedTo = common.MixAndMatchPurchaseItem
			case soloJazzDiscount:
				appliedTo = common.SoloJazzPurchaseItem
			case teamCompetitionDiscount:
				appliedTo = common.TeamCompetitionPurchaseItem
			case tshirtDiscount:
				appliedTo = common.TShirtPurchaseItem
			default:
				return nil, nil, fmt.Errorf("found unknown discount applied to %s", sd.AppliedTo)
			}
			result = append(result, &common.StoreDiscount{
				Name:      sd.Name,
				AppliedTo: appliedTo,
			})

			keys = append(keys, key[0].Encode())
		}
	}
	return keys, result, nil
}
