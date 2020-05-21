package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (s *Store) AddDiscount(ctx context.Context, discount *storage.Discount) error {

	stmt := fmt.Sprintf("WITH discount AS (INSERT INTO %s(%s) VALUES ($1) RETURNING %s AS id), discount_data(applied_to, name) AS ( VALUES ",
		discountBundleTable,
		discountBundleCodeCol,
		discountBundleIDCol)
	args := []interface{}{discount.Code}
	for i, d := range discount.Discounts {
		if i != 0 {
			stmt += ","
		}
		stmt += fmt.Sprintf("($%d::applied_to_type, $%d)", i*2+2, i*2+3)
		args = append(args, appliedToToEnum[d.AppliedTo], d.Name)
	}

	stmt += fmt.Sprintf(") INSERT INTO %s(%s, %s, %s) SELECT discount.id, discount_data.applied_to, discount_data.name FROM discount, discount_data", discountTable, discountFkCol, discountAppliedToCol, discountNameCol)

	_, err := s.pool.Exec(ctx, stmt, args...)
	if err != nil {
		var perr *pgconn.PgError
		if errors.As(err, &perr) {
			if perr.Code == pgerrcode.UniqueViolation && perr.ConstraintName == "discount_bundles_code_key" { // unique violation
				return storage.ErrDiscountExists{
					Code: discount.Code,
				}
			}
		}
		return fmt.Errorf("error inserting into database: %w", err)
	}
	return nil
}
