package authorizer

import "errors"

type Userinfo interface {
	UserID() string
	IsAllowed(permission string) bool
}

var Unauthenticated = errors.New("unauthenticated")
var Unauthorized = errors.New("unauthorized")
