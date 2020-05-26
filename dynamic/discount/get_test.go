package discount

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/sirupsen/logrus"
)

func compareSingleDiscount(t *testing.T, sd *Single, appliedTo storage.PurchaseItem, amount int) {
	t.Helper()
	if appliedTo != sd.AppliedTo {
		t.Fatalf("found unexpected appliedTo %v, expected %v", sd.AppliedTo, appliedTo)
	}

	amt, ok := sd.Amount.(common.DollarDiscount)
	if !ok {
		t.Fatalf("expected dollar discount, did not find it")
	}
	if int(amt) != amount {
		t.Fatalf("found unexpected discount amount %d, expected %d", amt, amount)
	}
}

func TestGetDiscount(t *testing.T) {
	expectedCode := "DJs rock"
	co := commontest.CommonCatalogObjects()

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	client := &commontest.MockSquareClient{
		ListCatalogFunc: commontest.ListCatalogFuncFromSlice(co.Catalog()),
	}

	store := &commontest.MockStore{
		GetDiscountFunc: func(ctx context.Context, code string) (*storage.Discount, error) {
			if code != expectedCode {
				t.Fatalf("found unexpected code in store call %s, expected %s", code, expectedCode)
			}

			return &storage.Discount{
				Code: code,
				Discounts: []*storage.SingleDiscount{
					{
						Name:      co.FullWeekendDiscountName,
						AppliedTo: storage.FullWeekendPurchaseItem,
					},
					{
						Name:      co.MixAndMatchDiscountName,
						AppliedTo: storage.MixAndMatchPurchaseItem,
					},
				},
			}, nil

		},
	}

	service := NewService(store, client, logger, &commontest.MockAuthorizer{})

	discount, err := service.Get(context.Background(), expectedCode)
	if err != nil {
		t.Fatalf("unexpected error in get discount call: %v", err)
	}
	if discount.Code != expectedCode {
		t.Fatalf("found unexpected code in discount %s, expected %s", discount.Code, expectedCode)
	}
	if len(discount.Discounts) != 2 {
		t.Fatalf("returned discount group contains too many single discounts %d, expected 2", len(discount.Discounts))
	}
	foundFullWeekend := false
	foundMixAndMatch := false
	for _, sd := range discount.Discounts {
		switch sd.Name {
		case co.FullWeekendDiscountName:
			foundFullWeekend = true
			compareSingleDiscount(t, sd, storage.FullWeekendPurchaseItem, co.FullWeekendDiscountAmount)
		case co.MixAndMatchDiscountName:
			foundMixAndMatch = true
			compareSingleDiscount(t, sd, storage.MixAndMatchPurchaseItem, co.MixAndMatchDiscountAmount)
		}
	}
	if !foundFullWeekend {
		t.Fatalf("did not find full weekend discount in result")
	}
	if !foundMixAndMatch {
		t.Fatalf("did not find mix and match discount in result")
	}
}

func TestGetNotExists(t *testing.T) {
	expectedCode := "DJs rock"

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	client := &commontest.MockSquareClient{
		ListCatalogFunc: func(context.Context, []square.CatalogObjectType) square.ListCatalogIterator {
			t.Fatalf("no need for square calls if discount does not exist")
			return nil
		},
	}

	store := &commontest.MockStore{
		GetDiscountFunc: func(ctx context.Context, code string) (*storage.Discount, error) {
			return nil, storage.ErrDiscountDoesNotExist{Code: code}
		},
	}

	service := NewService(store, client, logger, &commontest.MockAuthorizer{})

	_, err = service.Get(context.Background(), expectedCode)
	if err == nil {
		t.Fatalf("expected error, found none")
	}
	terr := storage.ErrDiscountDoesNotExist{}
	if !errors.As(err, &terr) {
		t.Fatalf("expected error to be of type \"discount does not exists\", found %v", err)
	}
	if terr.Code != expectedCode {
		t.Fatalf("expected code to be %s, got %s", expectedCode, terr.Code)
	}
}
