package authorizer

import "errors"

type Userinfo struct {
	UserID string
}

var Unauthenticated = errors.New("unauthenticated")
