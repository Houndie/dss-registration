package postgres

import (
	"context"
	"fmt"
)

func (s *Store) AddAdmin(ctx context.Context, userID string) error {
	_, err := s.pool.Exec(ctx, fmt.Sprintf("INSERT INTO %s(%s) VALUES ($1)", adminTable, adminUserIDCol), userID)
	if err != nil {
		return fmt.Errorf("error adding new admin to db: %w", err)
	}
	return nil
}
