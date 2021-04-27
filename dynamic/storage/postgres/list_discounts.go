package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

var listDiscountsStmt string

func init() {
	tmplStmt := `SELECT {{.SDAppliedToCol}}, {{.SDNameCol}}, {{.BundleCodeCol}}
		FROM {{.BundleTable}}, {{.SDTable}} 
		WHERE {{.BundleTable}}.{{.BundleIDCol}} = {{.SDTable}}.{{.SDFkCol}};`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing list discounts template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, discountConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing list discounts template: %v", err))
	}

	listDiscountsStmt = stmt.String()
}

func (s *Store) ListDiscounts(ctx context.Context) ([]*storage.Discount, error) {
	rows, err := s.pool.Query(ctx, listDiscountsStmt)
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
