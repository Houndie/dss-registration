package common

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
)

var ErrUnauthorized = errors.New("User is not authorized for this operation")

func IsAllowed(ctx context.Context, authorizer interface {
	GetUserinfo(context.Context, string) (authorizer.Userinfo, error)
}, token string, permission authorizer.Permission) error {
	userinfo, err := authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not authorize user: %w", err)
	}

	if !userinfo.IsAllowed(permission) {
		return ErrUnauthorized
	}
	return nil
}
