package postgres

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

const (
	discountBundleTable   = "discount_bundles"
	discountBundleCodeCol = "code"

	discountTable        = "discounts"
	discountFkCol        = "discount_bundle_id"
	discountNameCol      = "discount_name"
	discountAppliedToCol = "discount_applied_to"
)

var appliedToToEnum = map[storage.PurchaseItem]string{
	storage.FullWeekendPurchaseItem:     "Full Weekend",
	storage.DanceOnlyPurchaseItem:       "Dance Only",
	storage.MixAndMatchPurchaseItem:     "Mix And Match",
	storage.SoloJazzPurchaseItem:        "Solo Jazz",
	storage.TeamCompetitionPurchaseItem: "Team Competition",
	storage.TShirtPurchaseItem:          "TShirt",
}

func (s *Store) AddDiscount(ctx context.Context, discount *storage.Discount) error {
	var id uuid.UUID
	err := s.conn.QueryRow(fmt.Sprintf("INSERT INTO %s(%s) VALUES ($1);", discountBundleTable, discountBundleCodeCol), discount.Code).Scan(&id)
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES ", discountTable, discountFkCol, discountAppliedToCol, discountNameCol)
	args := []interface{}{}
	for i, d := range discount.Discounts {
		if i != 0 {
			stmt += ","
		}
		stmt += fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)
		args = append(args, appliedToToEnum[d.AppliedTo], d.Name)
	}

	_, err = s.conn.Exec(stmt, args...)
	if err != nil {
		return err
	}
	return nil
}
