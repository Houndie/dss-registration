package vaccine

import (
	"errors"
	"fmt"
)

type Info interface {
	isVaxApproval()
}

type VaxApproved struct{}
type VaxApprovalPending struct {
	URL string
}
type NoVaxProofSupplied struct{}

func (*VaxApproved) isVaxApproval()        {}
func (*VaxApprovalPending) isVaxApproval() {}
func (*NoVaxProofSupplied) isVaxApproval() {}

var ErrAlreadyApproved = errors.New("vaccine information already approved")

type ErrFileTooBig struct {
	Filesize int64
}

func (e ErrFileTooBig) Error() string {
	return fmt.Sprintf("filesize %d is too big", e.Filesize)
}
