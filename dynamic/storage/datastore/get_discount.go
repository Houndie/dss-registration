package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/getdiscount"
	"github.com/pkg/errors"
)

func (d *Datastore) GetDiscount(ctx context.Context, code string) ([]*getdiscount.StoreDiscount, error) {
	q := datastore.NewQuery(discountKind).Filter("Code =", code).Limit(1)
	discounts := []discountEntity{}
	_, err := d.client.GetAll(ctx, q, &discounts)
	if err != nil {
		return nil, errors.Wrapf(err, "error fetching discounts with code %s from datastore", code)
	}

	if len(discounts) == 0 {
		return nil, getdiscount.ErrDiscountDoesNotExist{
			Code: code,
		}
	}

	if len(discounts) > 1 {
		return nil, fmt.Errorf("somehow discovered %d discounts with code %s when only one was expected", len(discounts), code)
	}

	result := make([]*getdiscount.StoreDiscount, len(discounts[0].Discounts))

	for i, sd := range discounts[0].Discounts {
		var appliedTo common.DiscountTarget
		switch sd.AppliedTo {
		case fullWeekendDiscount:
			appliedTo = common.FullWeekendDiscountTarget
		case danceOnlyDiscount:
			appliedTo = common.DanceOnlyDiscountTarget
		case mixAndMatchDiscount:
			appliedTo = common.MixAndMatchDiscountTarget
		case soloJazzDiscount:
			appliedTo = common.SoloJazzDiscountTarget
		case teamCompetitionDiscount:
			appliedTo = common.TeamCompetitionDiscountTarget
		case tshirtDiscount:
			appliedTo = common.TShirtDiscountTarget
		default:
			return nil, fmt.Errorf("found unknown discount applied to %s", sd.AppliedTo)
		}
		result[i] = &getdiscount.StoreDiscount{
			Name:      sd.Name,
			AppliedTo: appliedTo,
		}
	}

	return result, nil
}
