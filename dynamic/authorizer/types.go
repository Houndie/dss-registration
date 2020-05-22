package authorizer

import "errors"

type Userinfo struct {
	UserId string
}

var Unauthenticated = errors.New("unauthenticated")
