package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
)

var isAdminStmt string

func init() {
	tmplStmt := "SELECT EXISTS(SELECT 1 FROM {{.Table}} WHERE {{.UserIDCol}} = $1);"

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing is admin template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, adminConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing is admin template: %v", err))
	}

	isAdminStmt = stmt.String()
}

func (s *Store) IsAdmin(ctx context.Context, userID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx, isAdminStmt, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error reading admin status from db: %w", err)
	}
	return exists, nil
}
