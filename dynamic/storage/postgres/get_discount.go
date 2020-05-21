package postgres

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Store) GetDiscount(ctx context.Context, code string) (*storage.Discount, error) {
	rows, err := s.pool.Query(ctx, fmt.Sprintf("SELECT %[1]s, %[2]s FROM %[3]s, %[4]s WHERE %[3]s.%[5]s = %[6]s AND %[7]s = $1", discountAppliedToCol, discountNameCol, discountBundleTable, discountTable, discountBundleIDCol, discountFkCol, discountBundleCodeCol), code)
	if err != nil {
		return nil, fmt.Errorf("error fetching discounts: %w", err)
	}
	defer rows.Close()

	result := &storage.Discount{
		Code:      code,
		Discounts: []*storage.SingleDiscount{},
	}

	found := false
	for rows.Next() {
		found = true
		var name string
		var appliedTo string
		err := rows.Scan(&appliedTo, &name)
		if err != nil {
			return nil, fmt.Errorf("error fetching discount result: %w", err)
		}
		result.Discounts = append(result.Discounts, &storage.SingleDiscount{
			Name:      name,
			AppliedTo: appliedToFromEnum[appliedTo],
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error closing rows: %w", err)
	}

	if !found {
		return nil, storage.ErrDiscountNotFound{Code: code}
	}
	return result, nil
}
