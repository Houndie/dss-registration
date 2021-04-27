package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestAddAdminAndIsAdmin(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)
	userID := "12345"
	defer func() {
		_, err := pool.Exec(context.Background(), "DELETE FROM "+adminTable)
		if err != nil {
			t.Fatalf("error cleaning up after test: %v", err)
		}
	}()
	if err := store.AddAdmin(context.Background(), userID); err != nil {
		t.Fatalf("error adding admin: %v", err)
	}

	isAdmin, err := store.IsAdmin(context.Background(), userID)
	if err != nil {
		t.Fatalf("error checking for admin existance: %v", err)
	}
	if !isAdmin {
		t.Fatalf("expected true, found false")
	}
}

func TestNotAdmin(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)
	userID := "12345"

	isAdmin, err := store.IsAdmin(context.Background(), userID)
	if err != nil {
		t.Fatalf("error checking for admin existance: %v", err)
	}
	if isAdmin {
		t.Fatalf("expected false, found true")
	}
}
