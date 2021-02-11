package postgres

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Store) ListDiscounts(ctx context.Context) ([]*storage.Discount, error) {
	rows, err := s.pool.Query(ctx, fmt.Sprintf("SELECT %[1]s, %[2]s, %[3]s FROM %[4]s, %[5]s WHERE %[4]s.%[6]s = %[7]s", discountAppliedToCol, discountNameCol, discountBundleCodeCol, discountBundleTable, discountTable, discountBundleIDCol, discountFkCol))
	if err != nil {
		return nil, fmt.Errorf("error fetching discounts: %w", err)
	}
	defer rows.Close()

	resultMap := map[string][]*storage.SingleDiscount{}
	for rows.Next() {
		var name string
		var appliedTo string
		var code string
		err := rows.Scan(&appliedTo, &name, &code)
		if err != nil {
			return nil, fmt.Errorf("error fetching discount result: %w", err)
		}

		found, ok := resultMap[code]
		if ok {
			resultMap[code] = append(found, &storage.SingleDiscount{
				Name:      name,
				AppliedTo: appliedToFromEnum[appliedTo],
			})
		} else {
			resultMap[code] = []*storage.SingleDiscount{
				{
					Name:      name,
					AppliedTo: appliedToFromEnum[appliedTo],
				},
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error closing rows: %w", err)
	}

	results := make([]*storage.Discount, len(resultMap))
	i := 0
	for code, singleDiscounts := range resultMap {
		results[i] = &storage.Discount{
			Code:      code,
			Discounts: singleDiscounts,
		}
		i++
	}
	return results, nil
}
