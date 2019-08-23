package listbyuser

import "time"

type Registration struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	Paid      bool
}

type StoreRegistration struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	OrderIds  []string
}
