package discount

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/sirupsen/logrus"
)

func TestAddDiscount(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	token := "some auth token"
	thisUserID := "123456"

	authorizer := &commontest.MockAuthorizer{
		UserinfoFunc: commontest.UserinfoFromIDCheck(t, token, thisUserID),
	}

	inDiscount := &Bundle{
		Code: "some code",
		Discounts: []*Single{
			{
				Name:      "DJ Discount",
				AppliedTo: storage.FullWeekendPurchaseItem,
			},
		},
	}

	store := &commontest.MockStore{
		AddDiscountFunc: func(ctx context.Context, discount *storage.Discount) error {
			if discount.Code != inDiscount.Code {
				t.Fatalf("expected discount code %s, found %s", inDiscount.Code, discount.Code)
			}
			for _, exp := range discount.Discounts {
				found := false
				for _, control := range inDiscount.Discounts {
					if exp.Name != control.Name {
						continue
					}

					found = true
					if exp.AppliedTo != control.AppliedTo {
						t.Fatalf("found applied to %s, expected %s", exp.AppliedTo, control.AppliedTo)
					}
					break
				}
				if !found {
					t.Fatalf("Did not find discount with name %s", exp.Name)
				}

			}
			return nil
		},
		IsAdminFunc: func(ctx context.Context, userID string) (bool, error) {
			if userID != thisUserID {
				t.Fatalf("expected user id %s, got %s", thisUserID, userID)
			}
			return true, nil
		},
	}

	service := NewService(store, &commontest.MockSquareClient{}, logger, authorizer)
	err = service.Add(context.Background(), token, inDiscount)
	if err != nil {
		t.Fatalf("unexpected error found in call to Add Discount: %v", err)
	}
}

func TestAddDiscountNotAuthorized(t *testing.T) {
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	token := "some auth token"
	thisUserID := "123456"

	authorizer := &commontest.MockAuthorizer{
		UserinfoFunc: commontest.UserinfoFromID(thisUserID),
	}

	inDiscount := &Bundle{
		Code: "some code",
		Discounts: []*Single{
			{
				Name:      "DJ Discount",
				AppliedTo: storage.FullWeekendPurchaseItem,
			},
		},
	}

	store := &commontest.MockStore{
		AddDiscountFunc: func(ctx context.Context, discount *storage.Discount) error {
			t.Fatalf("Discount added when non admin user given")
			return nil
		},
		IsAdminFunc: func(ctx context.Context, userID string) (bool, error) {
			return false, nil
		},
	}

	service := NewService(store, &commontest.MockSquareClient{}, logger, authorizer)
	err = service.Add(context.Background(), token, inDiscount)
	if err == nil {
		t.Fatalf("unexpected error, found none")
	}
	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("expected unauthorzed error, found: %v", err)
	}
}
