package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (d *Datastore) MarkOrderComplete(ctx context.Context, referenceId uuid.UUID, paymentId string) error {
	var orders []*orderEntity
	q := datastore.NewQuery(orderKind).Filter("ReferenceId = ", referenceId.String()).Limit(1)
	keys, err := d.client.GetAll(ctx, q, &orders)
	if err != nil {
		return errors.Wrap(err, "error retrieving order from datastore")
	}
	if len(orders) != 1 {
		return fmt.Errorf("found incorrect number of orders %d", len(orders))
	}
	if len(keys) != len(orders) {
		return fmt.Errorf("found incorrect number of keys %d, expected %d", len(keys), len(orders))
	}

	if orders[0].PaymentType != nonePayment {
		return storage.ErrAlreadyPaid{}
	}

	registrations[0].PaymentType = automaticPayment
	registrations[0].AutomaticPayment = paymentId
	_, err = d.client.Put(ctx, keys[0], registrations[0])
	return errors.Wrap(err, "error putting order back in datastore")
}
