package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

var getDiscountStmt string

func init() {
	tmplStmt := `SELECT {{.SDAppliedToCol}}, {{.SDNameCol}} 
		FROM {{.BundleTable}}, {{.SDTable}} 
		WHERE 
				{{.BundleTable}}.{{.BundleIDCol}} = {{.SDTable}}.{{.SDFkCol}} 
			AND 
				{{.BundleTable}}.{{.BundleCodeCol}} = $1`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing get discount template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, discountConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing get discount template: %v", err))
	}

	getDiscountStmt = stmt.String()
}

func (s *Store) GetDiscount(ctx context.Context, code string) (*storage.Discount, error) {
	rows, err := s.pool.Query(ctx, getDiscountStmt, code)
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
