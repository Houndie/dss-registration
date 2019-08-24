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
