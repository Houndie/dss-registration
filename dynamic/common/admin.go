package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("User is not authorized for this operation")

func IsAdmin(ctx context.Context, store interface {
	IsAdmin(context.Context, string) (bool, error)
}, authorizer interface {
	Userinfo(context.Context, string) (*authorizer.Userinfo, error)
}, logger *logrus.Logger, token string) error {
	logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := authorizer.Userinfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not authorize user: %w", err)
	}
	logger.Tracef("found user %s", userinfo.UserID)
	isAdmin, err := store.IsAdmin(ctx, userinfo.UserID)
	if err != nil {
		return fmt.Errorf("could not fetch admin status from store: %w", err)
	}

	if !isAdmin {
		return ErrUnauthorized
	}
	return nil
}
