package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (d *Datastore) GetDiscount(ctx context.Context, code string) (*storage.Discount, error) {
	q := datastore.NewQuery(discountKind).Filter("Code =", code).Limit(1)
	discounts := []*discountEntity{}
	_, err := d.client.GetAll(ctx, q, &discounts)
	if err != nil {
		return nil, fmt.Errorf("error fetching discounts with code %s from datastore: %w", code, err)
	}

	if len(discounts) == 0 {
		return nil, storage.ErrDiscountDoesNotExist{
			Code: code,
		}
	}

	if len(discounts) > 1 {
		return nil, fmt.Errorf("somehow discovered %d discounts with code %s when only one was expected", len(discounts), code)
	}

	result, err := fromDiscountEntity(discounts[0])
	if err != nil {
		return nil, fmt.Errorf("error converting discount entity to storage type")
	}

	return result, nil
}
