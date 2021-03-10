package auth0

import (
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/sirupsen/logrus"
)

type userinfo struct {
	userID      string
	permissions []authorizer.Permission
	logger      *logrus.Logger
}

func (u *userinfo) UserID() string {
	return u.userID
}

func (u *userinfo) IsAllowed(permission authorizer.Permission) bool {

	for _, listPermission := range u.permissions {
		if listPermission == permission {
			u.logger.Tracef("user %s granted permission %s", u.userID, permission)
			return true
		}
	}
	u.logger.Tracef("user %s not granted permission %s", u.userID, permission)
	return false
}
