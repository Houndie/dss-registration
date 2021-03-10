package discount

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/sirupsen/logrus"
)

func TestList(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	co := commontest.CommonCatalogObjects()
	client := &commontest.MockSquareClient{
		ListCatalogFunc: commontest.ListCatalogFuncFromSlice(co.Catalog()),
	}

	token := "some auth token"
	thisUserID := "123456"
	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, token, []authorizer.Permission{authorizer.ListDiscountsPermission}, thisUserID, []authorizer.Permission{authorizer.ListDiscountsPermission}),
	}
	code1 := "code1"
	code2 := "code2"
	store := &commontest.MockStore{
		ListDiscountsFunc: func(ctx context.Context) ([]*storage.Discount, error) {
			return []*storage.Discount{
				{
					Code: code1,
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
				}, {
					Code: code2,
					Discounts: []*storage.SingleDiscount{
						{
							Name:      co.MixAndMatchDiscountName,
							AppliedTo: storage.MixAndMatchPurchaseItem,
						},
					},
				},
			}, nil

		},
	}

	service := NewService(store, client, logger, authorizer)
	bundles, err := service.List(context.Background(), token)
	if err != nil {
		t.Fatalf("unexpected error from list: %v", err)
	}

	if len(bundles) != 2 {
		t.Fatalf("expected 2 bundles, found %d", len(bundles))
	}

	foundCode1 := false
	foundCode2 := false
	for _, bundle := range bundles {
		switch bundle.Code {
		case code1:
			foundCode1 = true
			if len(bundle.Discounts) != 2 {
				t.Fatalf("expected 2 discounts, found %v", len(bundle.Discounts))
			}

			foundFullWeekend := false
			foundMixAndMatch := false
			for _, d := range bundle.Discounts {
				switch d.Name {
				case co.FullWeekendDiscountName:
					foundFullWeekend = true
					if d.AppliedTo != storage.FullWeekendPurchaseItem {
						t.Fatalf("expected applied to %v, found %v", storage.FullWeekendPurchaseItem, d.AppliedTo)
					}
					amt, ok := d.Amount.(DollarDiscount)
					if !ok {
						t.Fatalf("discount is not a dollar discount")
					}
					if int(amt) != co.FullWeekendDiscountAmount {
						t.Fatalf("expected discount amount %d, found %d", co.FullWeekendDiscountAmount, amt)
					}
				case co.MixAndMatchDiscountName:
					foundMixAndMatch = true
					if d.AppliedTo != storage.MixAndMatchPurchaseItem {
						t.Fatalf("expected applied to %v, found %v", storage.MixAndMatchPurchaseItem, d.AppliedTo)
					}
					amt, ok := d.Amount.(DollarDiscount)
					if !ok {
						t.Fatalf("discount is not a dollar discount")
					}
					if int(amt) != co.MixAndMatchDiscountAmount {
						t.Fatalf("expected discount amount %d, found %d", co.MixAndMatchDiscountAmount, amt)
					}
				default:
					t.Fatalf("unknown discount name found: %s", d.Name)
				}
			}
			if !foundFullWeekend {
				t.Fatalf("did not find %s in discount bundle", co.FullWeekendDiscountName)
			}
			if !foundMixAndMatch {
				t.Fatalf("did not find %s in discount bundle", co.MixAndMatchDiscountName)
			}
		case code2:
			foundCode2 = true
			if len(bundle.Discounts) != 1 {
				t.Fatalf("expected 1 discount, found %v", len(bundle.Discounts))
			}
			if bundle.Discounts[0].AppliedTo != storage.MixAndMatchPurchaseItem {
				t.Fatalf("expected applied to %v, found %v", storage.MixAndMatchPurchaseItem, bundle.Discounts[0].AppliedTo)
			}
			if bundle.Discounts[0].Name != co.MixAndMatchDiscountName {
				t.Fatalf("expected name %v, found %v", co.MixAndMatchDiscountName, bundle.Discounts[0].Name)
			}
			amt, ok := bundle.Discounts[0].Amount.(DollarDiscount)
			if !ok {
				t.Fatalf("discount is not a dollar discount")
			}
			if int(amt) != co.MixAndMatchDiscountAmount {
				t.Fatalf("expected discount amount %d, found %d", co.MixAndMatchDiscountAmount, amt)
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

func TestListNone(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	client := &commontest.MockSquareClient{} // No square calls needed

	token := "some auth token"
	thisUserID := "123456"
	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(thisUserID, []authorizer.Permission{authorizer.ListDiscountsPermission}),
	}
	store := &commontest.MockStore{
		ListDiscountsFunc: func(ctx context.Context) ([]*storage.Discount, error) {
			return []*storage.Discount{}, nil

		},
	}

	service := NewService(store, client, logger, authorizer)
	bundles, err := service.List(context.Background(), token)
	if err != nil {
		t.Fatalf("unexpected error from list: %v", err)
	}

	if len(bundles) != 0 {
		t.Fatalf("expected 0 bundles, found %d", len(bundles))
	}
}

func TestListUnauthorized(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	client := &commontest.MockSquareClient{} // No square calls needed

	token := "some auth token"
	thisUserID := "123456"
	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(thisUserID, []authorizer.Permission{}),
	}
	store := &commontest.MockStore{}

	service := NewService(store, client, logger, authorizer)
	_, err = service.List(context.Background(), token)
	if err == nil {
		t.Fatalf("expected error from list, found none")
	}
	if !errors.Is(err, common.ErrUnauthorized) {
		t.Fatalf("expected unauthorized error, found: %v", err)
	}
}

func TestListUnauthenticated(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	client := &commontest.MockSquareClient{} // No square calls needed

	a := &commontest.MockAuthorizer{
		GetUserinfoFunc: func(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
			return nil, authorizer.Unauthenticated
		},
	}
	store := &commontest.MockStore{}

	service := NewService(store, client, logger, a)
	_, err = service.List(context.Background(), "some token")
	if err == nil {
		t.Fatalf("expected error from list, found none")
	}
	if !errors.Is(err, authorizer.Unauthenticated) {
		t.Fatalf("expected unauthorized error, found: %v", err)
	}
}
