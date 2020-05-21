package postgres

import (
	"context"
	"fmt"
)

func (s *Store) IsAdmin(ctx context.Context, userID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx, fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1);", adminTable, adminUserIDCol), userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error reading admin status from db: %w", err)
	}
	return exists, nil
}
