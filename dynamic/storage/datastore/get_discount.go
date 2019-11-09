package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/getdiscount"
	"github.com/pkg/errors"
)

func (d *Datastore) GetDiscount(ctx context.Context, code string) ([]*common.StoreDiscount, error) {
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

	result := make([]*common.StoreDiscount, len(discounts[0].Discounts))

	for i, sd := range discounts[0].Discounts {
		appliedTo, err := parseAppliedTo(sd.AppliedTo)
		if err != nil {
			return nil, errors.Wrap(nil, "found unknown error appliedto from store")
		}
		result[i] = &common.StoreDiscount{
			Name:      sd.Name,
			AppliedTo: appliedTo,
		}
	}

	return result, nil
}

func parseAppliedTo(datastoreVal string) (common.PurchaseItem, error) {
	var appliedTo common.PurchaseItem
	switch datastoreVal {
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
		return "", fmt.Errorf("found unknown discount applied to %s", datastoreVal)
	}
	return appliedTo, nil
}
