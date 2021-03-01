package postgres

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
)

var updateDiscountStmt1 string
var updateDiscountTmpl2 *template.Template

func init() {
	tmplStmt1 :=
		`DELETE 
			FROM {{.SDTable}} 
			USING {{.BundleTable}}
			WHERE 
					{{.SDTable}}.{{.SDFkCol}} = {{.BundleTable}}.{{.BundleIDCol}}
				AND
					{{.BundleCodeCol}} = $1`

	tmplStmt2 :=
		`WITH discount AS (
			UPDATE {{.C.BundleTable}}
			SET 
				{{.C.BundleCodeCol}} = $1
			WHERE
				{{.C.BundleCodeCol}} = $2
			RETURNING
				{{.C.BundleIDCol}}
		), discount_data(applied_to, name) AS (
			VALUES {{ range $idx, $elem := .V }} 
				{{if $idx}},{{end}}(${{value1 $idx}}::applied_to_type, ${{value2 $idx}}) 
			{{ end }}
		) INSERT INTO {{ .C.SDTable }} ({{ .C.SDFkCol }}, {{ .C.SDAppliedToCol }}, {{ .C.SDNameCol }}) 
			SELECT discount.id, discount_data.applied_to, discount_data.name 
			FROM discount, discount_data;`

	tmpl1, err := template.New("tmpl").Parse(tmplStmt1)
	if err != nil {
		panic(fmt.Sprintf("error parsing first update discount template: %v", err))
	}

	updateDiscountTmpl2, err = template.New("tmpl").Funcs(template.FuncMap{
		"value1": func(idx int) int {
			return idx*2 + 3
		},
		"value2": func(idx int) int {
			return idx*2 + 4
		},
	}).Parse(tmplStmt2)
	if err != nil {
		panic(fmt.Sprintf("error parsing second update discount template: %v", err))
	}

	stmt1 := &bytes.Buffer{}
	err = tmpl1.Execute(stmt1, discountConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing first update discount template: %v", err))
	}
	updateDiscountStmt1 = stmt1.String()
}

func (s *Store) UpdateDiscount(ctx context.Context, oldCode string, newDiscount *storage.Discount) error {
	stmt2 := &bytes.Buffer{}
	vals := struct {
		C *discountConstsTypes
		V []*storage.SingleDiscount
	}{
		C: discountConsts,
		V: newDiscount.Discounts,
	}
	err := updateDiscountTmpl2.Execute(stmt2, vals)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	args2 := []interface{}{newDiscount.Code, oldCode}
	for _, d := range newDiscount.Discounts {
		args2 = append(args2, appliedToToEnum[d.AppliedTo], d.Name)
	}

	tx, err := s.pool.Begin(ctx)
	defer tx.Rollback(ctx)

	queries := &pgx.Batch{}
	queries.Queue(updateDiscountStmt1, oldCode)
	queries.Queue(stmt2.String(), args2...)

	results := tx.SendBatch(ctx, queries)
	defer results.Close()
	res, err := results.Exec() // Delete
	if err != nil {
		return fmt.Errorf("error deleting single discount rows: %w", err)
	}
	if res.RowsAffected() == 0 {
		return storage.ErrDiscountNotFound{Code: oldCode}
	}

	_, err = results.Exec() // Update/Insert
	if err != nil {
		var perr *pgconn.PgError
		if errors.As(err, &perr) {
			if perr.Code == pgerrcode.UniqueViolation && perr.ConstraintName == "discount_bundles_code_key" { // unique violation
				return storage.ErrDiscountExists{
					Code: newDiscount.Code,
				}
			}
		}
		return fmt.Errorf("error updating discount bundle, and inserting new single discounts: %w", err)
	}

	err = results.Close()
	if err != nil {
		return fmt.Errorf("error closing sql results: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil
}
