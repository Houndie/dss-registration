package storage

import "fmt"

type ErrAlreadyPaid struct{}

func (ErrAlreadyPaid) Error() string {
	return "registration already marked as paid"
}

type ErrNotFound struct {
	Key string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("registration item with key %s not found in datastore", e.Key)
}

type ErrVolunteerExists struct {
	UserId string
}

func (e ErrVolunteerExists) Error() string {
	return fmt.Sprintf("volunteer submission for userid %s already exists", e.UserId)
}

type ErrDiscountExists struct {
	Code string
}

func (e ErrDiscountExists) Error() string {
	return fmt.Sprintf("discount with code %s already exists", e.Code)
}

type ErrDiscountNotFound struct {
	Code string
}

func (e ErrDiscountNotFound) Error() string {
	return fmt.Sprintf("discount with code %s not found", e.Code)
}

type ErrNoRegistrationForID struct {
	ID string
}

func (e ErrNoRegistrationForID) Error() string {
	return fmt.Sprintf("no registration found for id %s", e.ID)
}
