package auth0

import (
	"github.com/sirupsen/logrus"
)

type userinfo struct {
	userID      string
	permissions []string
	logger      *logrus.Logger
}

func (u *userinfo) UserID() string {
	return u.userID
}

func (u *userinfo) IsAllowed(permission string) bool {

	for _, listPermission := range u.permissions {
		if listPermission == permission {
			u.logger.Tracef("user %s granted permission %s", u.userID, permission)
			return true
		}
	}
	u.logger.Tracef("user %s not granted permission %s", u.userID, permission)
	return false
}
