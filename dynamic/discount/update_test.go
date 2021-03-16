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
	"github.com/Houndie/square-go"
	"github.com/sirupsen/logrus"
)

func TestUpdateDiscount(t *testing.T) {
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
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, token, []authorizer.Permission{authorizer.EditDiscountPermission}, thisUserID, []authorizer.Permission{authorizer.EditDiscountPermission}),
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

	inCode := "some old code"

	store := &commontest.MockStore{
		UpdateDiscountFunc: func(ctx context.Context, oldCode string, newDiscount *storage.Discount) error {
			if newDiscount.Code != inDiscount.Code {
				t.Fatalf("expected discount code %s, found %s", inDiscount.Code, newDiscount.Code)
			}
			if oldCode != inCode {
				t.Fatalf("expected old code %s, found %s", inCode, oldCode)
			}
			for _, exp := range newDiscount.Discounts {
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
	}

	service := NewService(store, &square.Client{}, logger, authorizer)
	err = service.Update(context.Background(), token, inCode, inDiscount)
	if err != nil {
		t.Fatalf("unexpected error found in call to Add Discount: %v", err)
	}
}

func TestUpdateDiscountNotAuthorized(t *testing.T) {
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
		GetUserinfoFunc: commontest.UserinfoFromID(thisUserID, []authorizer.Permission{}),
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

	inCode := "some old code"

	store := &commontest.MockStore{}

	service := NewService(store, &square.Client{}, logger, authorizer)
	err = service.Update(context.Background(), token, inCode, inDiscount)
	if err == nil {
		t.Fatalf("unexpected error, found none")
	}
	if !errors.Is(err, common.ErrUnauthorized) {
		t.Fatalf("expected unauthorzed error, found: %v", err)
	}
}
