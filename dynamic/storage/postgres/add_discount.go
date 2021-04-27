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
)

var addDiscountTmpl *template.Template

func init() {
	tmplStmt :=
		`WITH discount AS (
			INSERT INTO 
				{{ .C.BundleTable }} ({{ .C.BundleCodeCol }}) 
			VALUES ($1) 
			RETURNING {{ .C.BundleIDCol }} AS id
		), discount_data(applied_to, name) AS (
			VALUES {{ range $idx, $elem := .V }} 
				{{if $idx}},{{end}}(${{value1 $idx}}::applied_to_type, ${{value2 $idx}}) 
			{{ end }}
		) INSERT INTO {{ .C.SDTable }} ({{ .C.SDFkCol }}, {{ .C.SDAppliedToCol }}, {{ .C.SDNameCol }}) 
			SELECT discount.id, discount_data.applied_to, discount_data.name 
			FROM discount, discount_data;`

	var err error
	addDiscountTmpl, err = template.New("tmpl").Funcs(template.FuncMap{
		"value1": func(idx int) int {
			return idx*2 + 2
		},
		"value2": func(idx int) int {
			return idx*2 + 3
		},
	}).Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing template: %v", err))
	}

}

func (s *Store) AddDiscount(ctx context.Context, discount *storage.Discount) error {
	stmt := &bytes.Buffer{}
	vals := struct {
		C *discountConstsTypes
		V []*storage.SingleDiscount
	}{
		C: discountConsts,
		V: discount.Discounts,
	}
	err := addDiscountTmpl.Execute(stmt, vals)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	args := []interface{}{discount.Code}
	for _, d := range discount.Discounts {
		args = append(args, appliedToToEnum[d.AppliedTo], d.Name)
	}

	_, err = s.pool.Exec(ctx, stmt.String(), args...)
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
