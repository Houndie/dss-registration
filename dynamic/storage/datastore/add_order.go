package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func (s *Datastore) AddOrder(ctx context.Context, registrationId string, o *StoreOrder) error {
	order := &orderEntity{
		referenceId: o.ReferenceId.String(),
		orderId:     o.OrderId,
		PaymentType: nonePayment,
	}

	orderKey := datastore.IncompleteKey(orderKind, registrationId)
	_, err := datastore.Put(ctx, orderKey, order)
	if err != nil {
		return errors.Wrap(err, "error adding order to datastore")
	}
	return nil

}
