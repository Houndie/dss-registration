package discount

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/sirupsen/logrus"
)

func TestDeleteDiscount(t *testing.T) {
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
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, token, []authorizer.Permission{authorizer.DeleteDiscountPermission}, thisUserID, []authorizer.Permission{authorizer.DeleteDiscountPermission}),
	}

	inCode := "some code"

	store := &commontest.MockStore{
		DeleteDiscountFunc: func(ctx context.Context, code string) error {
			if code != inCode {
				t.Fatalf("expected code %s, found %s", inCode, code)
			}
			return nil
		},
	}

	service := NewService(store, &square.Client{}, logger, authorizer)
	err = service.Delete(context.Background(), token, inCode)
	if err != nil {
		t.Fatalf("unexpected error found in call to Add Discount: %v", err)
	}
}

func TestDeleteDiscountNotAuthorized(t *testing.T) {
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

	inCode := "some code"

	store := &commontest.MockStore{}

	service := NewService(store, &square.Client{}, logger, authorizer)
	err = service.Delete(context.Background(), token, inCode)
	if err == nil {
		t.Fatalf("unexpected error, found none")
	}
	if !errors.Is(err, common.ErrUnauthorized) {
		t.Fatalf("expected unauthorzed error, found: %v", err)
	}
}
