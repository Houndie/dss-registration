package commontest

import (
	"context"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
)

type MockAuthorizer struct {
	GetUserinfoFunc func(ctx context.Context, accessToken string) (authorizer.Userinfo, error)
}

func (m *MockAuthorizer) GetUserinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
	return m.GetUserinfoFunc(ctx, accessToken)
}

type MockUserinfo struct {
	UserIDFunc    func() string
	IsAllowedFunc func(permission authorizer.Permission) bool
}

func (m *MockUserinfo) UserID() string {
	return m.UserIDFunc()
}

func (m *MockUserinfo) IsAllowed(permission authorizer.Permission) bool {
	return m.IsAllowedFunc(permission)
}

func UserinfoFromIDCheck(t *testing.T, expectedToken string, expectedPermissions []authorizer.Permission, ID string, permissions []authorizer.Permission) func(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
	return func(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
		if accessToken == "" {
			t.Fatalf("authorizer called but no accessToken provided")
		}
		if accessToken != expectedToken {
			t.Fatalf("found unexpected access token %s, expected %s", accessToken, expectedToken)
		}
		return &MockUserinfo{
			UserIDFunc: func() string { return ID },
			IsAllowedFunc: func(permission authorizer.Permission) bool {
				found := false
				for _, p := range expectedPermissions {
					if permission == p {
						found = true
					}
				}
				if !found {
					t.Fatalf("unexpected permission found: %v", permission)
				}

				for _, p := range permissions {
					if permission == p {
						return true
					}
				}
				return false
			},
		}, nil
	}
}

func UserinfoFromID(ID string, permissions []authorizer.Permission) func(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
	return func(ctx context.Context, accessToken string) (authorizer.Userinfo, error) {
		return &MockUserinfo{
			UserIDFunc: func() string { return ID },
			IsAllowedFunc: func(permission authorizer.Permission) bool {
				for _, p := range permissions {
					if permission == p {
						return true
					}
				}
				return false
			},
		}, nil
	}
}
