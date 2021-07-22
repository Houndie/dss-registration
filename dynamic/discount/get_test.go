package discount

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/sirupsen/logrus"
)

func compareSingleDiscount(t *testing.T, sd *common.Discount, appliedTo storage.PurchaseItem, amount int) {
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
	co := commontest.CommonCatalogObjects()
	expectedCode := co.FullWeekendDiscountName

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	service := NewService(logger, co.SquareData())

	discount, err := service.Get(context.Background(), expectedCode)
	if err != nil {
		t.Fatalf("unexpected error in get discount call: %v", err)
	}
	if len(discount) != 1 {
		t.Fatalf("returned discount group contains too many single discounts %d, expected 1", len(discount))
	}
	compareSingleDiscount(t, discount[0], storage.FullWeekendPurchaseItem, co.FullWeekendDiscountAmount)
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

	service := NewService(logger, commontest.CommonCatalogObjects().SquareData())

	_, err = service.Get(context.Background(), expectedCode)
	if err == nil {
		t.Fatalf("expected error, found none")
	}
	terr := storage.ErrDiscountNotFound{}
	if !errors.As(err, &terr) {
		t.Fatalf("expected error to be of type \"discount does not exists\", found %v", err)
	}
	if terr.Code != expectedCode {
		t.Fatalf("expected code to be %s, got %s", expectedCode, terr.Code)
	}
}
