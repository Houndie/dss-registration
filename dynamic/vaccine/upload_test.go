package vaccine

import (
	"context"
	"errors"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/object/aws"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

var testPermissionConfig = &PermissionConfig{
	Approve: "approve:vaccine",
	Get:     "get:vaccine",
	Upload:  "upload:vaccine",
}

func TestUploadVaxImage(t *testing.T) {
	logger := logrus.New()

	testID := "testid"
	testUserID := "test_user_id"
	testToken := "token"

	store := &commontest.MockStore{
		GetRegistrationFunc: func(ctx context.Context, id string) (*storage.Registration, error) {
			if id != testID {
				t.Fatalf("unexpected id %s", id)
			}

			return &storage.Registration{
				ID:     id,
				UserID: testUserID,
			}, nil
		},

		GetVaccineFunc: func(ctx context.Context, id string) (bool, error) {
			if id != testID {
				t.Fatalf("unexpected id %s", id)
			}

			return false, nil
		},
	}

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(testUserID, []string{}),
	}

	objectClient, err := aws.NewObjectClient("access", "secret", "region", "bucket")
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(logger, authorizer, store, objectClient, nil, testPermissionConfig)
	_, err = service.Upload(context.Background(), testToken, 1234, testID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadVaxImageNotMyRegistration(t *testing.T) {
	logger := logrus.New()

	testID := "testid"
	testUserID := "test_user_id"
	testToken := "token"

	store := &commontest.MockStore{
		GetRegistrationFunc: func(ctx context.Context, id string) (*storage.Registration, error) {
			if id != testID {
				t.Fatalf("unexpected id %s", id)
			}

			return &storage.Registration{
				ID:     id,
				UserID: "bad user id",
			}, nil
		},
	}

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(testUserID, []string{}),
	}

	objectClient, err := aws.NewObjectClient("access", "secret", "region", "bucket")
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(logger, authorizer, store, objectClient, nil, testPermissionConfig)
	_, err = service.Upload(context.Background(), testToken, 1234, testID)
	if err == nil {
		t.Fatal("expected error, found none")
	}

	var e storage.ErrNotFound
	if !errors.As(err, &e) {
		t.Fatal(err)
	}

	if e.Key != testID {
		t.Fatalf("unxpected id %s", e.Key)
	}
}

func TestUploadVaxImageAlreadyApproved(t *testing.T) {
	logger := logrus.New()

	testID := "testid"
	testUserID := "test_user_id"
	testToken := "token"

	store := &commontest.MockStore{
		GetRegistrationFunc: func(ctx context.Context, id string) (*storage.Registration, error) {
			if id != testID {
				t.Fatalf("unexpected id %s", id)
			}

			return &storage.Registration{
				ID:     id,
				UserID: testUserID,
			}, nil
		},

		GetVaccineFunc: func(ctx context.Context, id string) (bool, error) {
			if id != testID {
				t.Fatalf("unexpected id %s", id)
			}

			return true, nil
		},
	}

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(testUserID, []string{}),
	}

	service := NewService(logger, authorizer, store, nil, nil, testPermissionConfig)
	_, err := service.Upload(context.Background(), testToken, 1234, testID)
	if err == nil {
		t.Fatal("expected error")
	}
}
