package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

var deleteDiscountStmt string

func init() {
	tmplStmt := `DELETE FROM {{.BundleTable}} WHERE {{.BundleCodeCol}} = $1;`
	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing delete discount template: %v", err))
	}
	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, discountConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing delete discount template: %v", err))
	}
	deleteDiscountStmt = stmt.String()
}

func (s *Store) DeleteDiscount(ctx context.Context, code string) error {
	res, err := s.pool.Exec(ctx, deleteDiscountStmt, code)
	if err != nil {
		return fmt.Errorf("error deleting discount: %w", err)
	}
	if res.RowsAffected() == 0 {
		return storage.ErrDiscountNotFound{Code: code}
	}
	return nil
}
