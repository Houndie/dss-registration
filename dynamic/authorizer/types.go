package authorizer

import "errors"

type Userinfo interface {
	UserID() string
	IsAllowed(permission Permission) bool
}

var Unauthenticated = errors.New("unauthenticated")

type Permission string

const (
	ListDiscountsPermission  Permission = "list:discounts"
	AddDiscountPermission    Permission = "add:discounts"
	DeleteDiscountPermission Permission = "delete:discounts"
	EditDiscountPermission   Permission = "edit:discounts"
)

func (p Permission) IsValid() bool {
	switch p {
	case ListDiscountsPermission,
		AddDiscountPermission,
		DeleteDiscountPermission,
		EditDiscountPermission:
		return true
	}
	return false

}
