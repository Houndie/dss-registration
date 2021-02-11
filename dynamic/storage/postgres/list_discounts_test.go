package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestListDiscounts(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	store := NewStore(pool)
	code1 := "code1"
	code2 := "code2"
	name1 := "name1"
	name2 := "name2"
	name3 := "name3"
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
	rows, err := pool.Query(context.Background(), fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1), ($2) RETURNING %s;", discountBundleTable, discountBundleCodeCol, discountBundleIDCol), code1, code2)
	if err != nil {
		t.Fatalf("error inserting discount bundles for test: %v", err)
	}
	defer rows.Close()
	ids := []string{}
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			t.Fatalf("error getting id from bundle: %v", err)
		}

		ids = append(ids, id)
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3), ($1, $4, $5), ($6, $7, $8);", discountTable, discountFkCol, discountAppliedToCol, discountNameCol), ids[0], appliedToToEnum[storage.FullWeekendPurchaseItem], name1, appliedToToEnum[storage.DanceOnlyPurchaseItem], name2, ids[1], appliedToToEnum[storage.MixAndMatchPurchaseItem], name3)
	if err != nil {
		t.Fatalf("error inserting discounts for test: %v", err)
	}

	discounts, err := store.ListDiscounts(context.Background())
	if err != nil {
		t.Fatalf("error listing discounts from store: %v", err)
	}
	if len(discounts) != 2 {
		t.Fatalf("expected 2 discounts, found %v", len(discounts))
	}

	foundCode1 := false
	foundCode2 := false
	for _, bundle := range discounts {
		switch bundle.Code {
		case code1:
			foundCode1 = true
			if len(bundle.Discounts) != 2 {
				t.Fatalf("expected 2 discounts, found %v", len(bundle.Discounts))
			}

			foundName1 := false
			foundName2 := false
			for _, d := range bundle.Discounts {
				switch d.Name {
				case name1:
					foundName1 = true
					if d.AppliedTo != storage.FullWeekendPurchaseItem {
						t.Fatalf("expected applied to %v, found %v", storage.FullWeekendPurchaseItem, d.AppliedTo)
					}
				case name2:
					foundName2 = true
					if d.AppliedTo != storage.DanceOnlyPurchaseItem {
						t.Fatalf("expected applied to %v, found %v", storage.DanceOnlyPurchaseItem, d.AppliedTo)
					}
				default:
					t.Fatalf("unknown discount name found: %s", d.Name)
				}
			}
			if !foundName1 {
				t.Fatalf("did not find %s in discount bundle", name1)
			}
			if !foundName2 {
				t.Fatalf("did not find %s in discount bundle", name2)
			}
		case code2:
			foundCode2 = true
			if len(bundle.Discounts) != 1 {
				t.Fatalf("expected 1 discount, found %v", len(bundle.Discounts))
			}
			if bundle.Discounts[0].AppliedTo != storage.MixAndMatchPurchaseItem {
				t.Fatalf("expected applied to %v, found %v", storage.MixAndMatchPurchaseItem, bundle.Discounts[0].AppliedTo)
			}
			if bundle.Discounts[0].Name != name3 {
				t.Fatalf("expected name %v, found %v", name3, bundle.Discounts[0].Name)
			}
		}
	}
	if !foundCode1 {
		t.Fatalf("did not find %s in discounts", code1)
	}
	if !foundCode2 {
		t.Fatalf("did not find %s in discounts", code2)
	}
}

func TestListDiscountsNone(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	store := NewStore(pool)
	bundles, err := store.ListDiscounts(context.Background())
	if err != nil {
		t.Fatalf("error listing discounts: %v", err)
	}
	if len(bundles) != 0 {
		t.Fatalf("found %d bundles, expected none", len(bundles))
	}
}
