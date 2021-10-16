package auth0

import (
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/lestrrat-go/jwx/jwt"
)

func (a *Authorizer) GetUserinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
	a.logger.Tracef("parsing jwt: %s", accessToken)
	jwks, err := a.jwks.get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching jwks: %w", err)
	}
	token, err := jwt.ParseString(accessToken, jwt.WithKeySet(jwks), jwt.WithValidate(true), jwt.WithAcceptableSkew(1*time.Minute))
	if err != nil {
		return nil, fmt.Errorf("error parsing jwt: %v", err)
	}

	permissions, ok := token.Get("permissions")
	if !ok {
		a.logger.Tracef("access token found for user %s", token.Subject())
		return &userinfo{
			userID:      token.Subject(),
			permissions: []authorizer.Permission{},
			logger:      a.logger,
		}, nil
	}
	permissionsList, ok := permissions.([]interface{})
	if !ok {
		return nil, errors.New("permissions list not of correct type")
	}
	p := make([]authorizer.Permission, len(permissionsList))
	for i, perm := range permissionsList {
		str, ok := perm.(string)
		if !ok {
			return nil, errors.New("permission not of correct type")
		}
		p[i] = authorizer.Permission(str)
		if !p[i].IsValid() {
			return nil, fmt.Errorf("unknown permission assigned to user")
		}
	}
	a.logger.Tracef("access token found for user %s", token.Subject())
	return &userinfo{
		userID:      token.Subject(),
		permissions: p,
		logger:      a.logger,
	}, nil
}
