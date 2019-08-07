package storage

type ErrAlreadyPaid struct{}

func (ErrAlreadyPaid) Error() string {
	return "registration already marked as paid"
}
