package postgres

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestGetDiscountDoesntExist(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_PG_URL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	store := NewStore(pool)
	expectedCode := "code"
	_, err = store.GetDiscount(context.Background(), expectedCode)
	if err == nil {
		t.Fatalf("expected error fetching non-existant discount, found none")
	}

	serr := storage.ErrDiscountNotFound{}
	if !errors.As(err, &serr) {
		t.Fatalf("expected discount exists error, found %v", err)
	}

	if serr.Code != expectedCode {
		t.Fatalf("expected discount code %s, found %s", expectedCode, serr.Code)
	}
}
