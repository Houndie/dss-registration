package registration

import (
	"context"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
)

type mockAuthorizer struct {
	UserinfoFunc func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

func (m *mockAuthorizer) Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
	return m.UserinfoFunc(ctx, accessToken)
}

func UserinfoFromIDCheck(t *testing.T, expectedToken, ID string) func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
	return func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
		if accessToken == "" {
			t.Fatalf("authorizer called but no accessToken provided")
		}
		if accessToken != expectedToken {
			t.Fatalf("found unexpected access token %s, expected %s", accessToken, expectedToken)
		}
		return &authorizer.Userinfo{
			UserId: ID,
		}, nil
	}
}

func UserinfoFromID(ID string) func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
	return func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
		return &authorizer.Userinfo{
			UserId: ID,
		}, nil
	}
}
