package registration

import (
	"context"
	"errors"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/object/aws"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

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
	}

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID(testUserID, []authorizer.Permission{}),
	}

	objectClient, err := aws.NewObjectClient("access", "secret", "region", "bucket")
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(true, true, logger, nil, nil, authorizer, store, nil, objectClient)
	_, err = service.UploadVaxImage(context.Background(), testToken, 1234, testID)
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
		GetUserinfoFunc: commontest.UserinfoFromID(testUserID, []authorizer.Permission{}),
	}

	objectClient, err := aws.NewObjectClient("access", "secret", "region", "bucket")
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(true, true, logger, nil, nil, authorizer, store, nil, objectClient)
	_, err = service.UploadVaxImage(context.Background(), testToken, 1234, testID)
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
