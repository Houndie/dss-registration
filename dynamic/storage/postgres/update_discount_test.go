package postgres

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestUpdateDiscount(t *testing.T) {
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

	discount2 := &storage.Discount{
		Code: "code2",
		Discounts: []*storage.SingleDiscount{
			{
				AppliedTo: storage.FullWeekendPurchaseItem,
				Name:      "full weekend name 2",
			},
			{
				AppliedTo: storage.DanceOnlyPurchaseItem,
				Name:      "dance only name 2",
			},
			{
				AppliedTo: storage.MixAndMatchPurchaseItem,
				Name:      "mix and match name 2",
			},
			{
				AppliedTo: storage.SoloJazzPurchaseItem,
				Name:      "solo jazz name 2",
			},
			{
				AppliedTo: storage.TeamCompetitionPurchaseItem,
				Name:      "team competition name 2",
			},
			{
				AppliedTo: storage.TShirtPurchaseItem,
				Name:      "t-shirt name 2",
			},
		},
	}

	err = store.UpdateDiscount(context.Background(), discount.Code, discount2)
	if err != nil {
		t.Fatalf("error updating discount: %v", err)
	}

	testDiscount, err := store.GetDiscount(context.Background(), discount2.Code)
	if err != nil {
		t.Fatalf("error getting discount: %v", err)
	}
	if discount2.Code != testDiscount.Code {
		t.Fatalf("found discount code %s, expected %s", testDiscount.Code, discount2.Code)
	}

	if len(discount2.Discounts) != len(testDiscount.Discounts) {
		t.Fatalf("found %d individual discounts, expected %d", len(testDiscount.Discounts), len(discount2.Discounts))
	}

	for _, tsd := range testDiscount.Discounts {
		found := false
		for _, sd := range discount2.Discounts {
			if sd.Name == tsd.Name {
				found = true

				if sd.AppliedTo != tsd.AppliedTo {
					t.Fatalf("found applied to %v, expected %v", tsd.AppliedTo, sd.AppliedTo)
				}
				break
			}

		}
		if !found {
			t.Fatalf("unable to find single discount with name %s", tsd.Name)
		}
	}
}

func TestUpdateDiscountDuplicate(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	discount := &storage.Discount{
		Code: "code",
		Discounts: []*storage.SingleDiscount{
			{
				AppliedTo: storage.FullWeekendPurchaseItem,
				Name:      "full weekend name",
			},
		},
	}
	discount2 := &storage.Discount{
		Code: "code2",
		Discounts: []*storage.SingleDiscount{
			{
				AppliedTo: storage.FullWeekendPurchaseItem,
				Name:      "full weekend name",
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
	err = store.AddDiscount(context.Background(), discount2)
	if err != nil {
		t.Fatalf("error adding discount: %v", err)
	}
	err = store.UpdateDiscount(context.Background(), discount2.Code, discount)
	if err == nil {
		t.Fatalf("expected error updating discount with existing code, found none")
	}

	serr := storage.ErrDiscountExists{}
	if !errors.As(err, &serr) {
		t.Fatalf("expected discount exists error, found %v", err)
	}

	if serr.Code != discount.Code {
		t.Fatalf("expected discount code %s, found %s", discount.Code, serr.Code)
	}
}

func TestUpdateDiscountDoesntExist(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	store := NewStore(pool)
	expectedCode := "code"
	discount := &storage.Discount{
		Code: "new code",
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
	err = store.UpdateDiscount(context.Background(), expectedCode, discount)
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
