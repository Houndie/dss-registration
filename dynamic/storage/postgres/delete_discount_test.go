package postgres

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestDeleteDiscount(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()
	discount := &storage.Discount{
		Code: "code",
		Discounts: []*storage.SingleDiscount{
			{
				AppliedTo: storage.FullWeekendPurchaseItem,
				Name:      "full weekend name",
			},
			{
				AppliedTo: storage.DanceOnlyPurchaseItem,
				Name:      "dance only name",
			},
			{
				AppliedTo: storage.MixAndMatchPurchaseItem,
				Name:      "mix and match name",
			},
			{
				AppliedTo: storage.SoloJazzPurchaseItem,
				Name:      "solo jazz name",
			},
			{
				AppliedTo: storage.TeamCompetitionPurchaseItem,
				Name:      "team competition name",
			},
			{
				AppliedTo: storage.TShirtPurchaseItem,
				Name:      "t-shirt name",
			},
		},
	}
	store := NewStore(pool)
	defer func() {
		_, err := pool.Exec(context.Background(), "DELETE FROM "+discountBundleTable)
		if err != nil {
			t.Fatalf("error cleaning up after test: %v", err)
		}
		_, err = pool.Exec(context.Background(), "DELETE FROM "+discountTable)
		if err != nil {
			t.Fatalf("error cleaning up after test: %v", err)
		}
	}()
	err = store.AddDiscount(context.Background(), discount)
	if err != nil {
		t.Fatalf("error adding discount: %v", err)
	}

	err = store.DeleteDiscount(context.Background(), discount.Code)
	if err != nil {
		t.Fatalf("error deleting discount: %v", err)
	}

	_, err = store.GetDiscount(context.Background(), discount.Code)
	if err == nil {
		t.Fatalf("expected error getting deleted discount, found none")
	}

	e := storage.ErrDiscountNotFound{}
	if !errors.As(err, &e) {
		t.Fatalf("found error different than expected: %v", err)
	}

	if e.Code != discount.Code {
		t.Fatalf("found code %s, expected %s", e.Code, discount.Code)
	}
}
