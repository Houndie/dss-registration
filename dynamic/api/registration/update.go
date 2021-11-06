package registration

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

func (s *Server) Update(ctx context.Context, req *pb.RegistrationUpdateReq) (*pb.RegistrationUpdateRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	info, err := fromProtoc(req.Registration)
	if err != nil {
		var parseErr errParseField
		if errors.As(err, &parseErr) {
			switch parseErr.field {
			case parseFieldFailureCreatedAt:
				return nil, twirp.InvalidArgumentError("registration.created_at", err.Error())
			case parseFieldFailureTier:
				return nil, twirp.InvalidArgumentError("registration.full_weekend_pass.tier", err.Error())
			case parseFieldFailureLevel:
				return nil, twirp.InvalidArgumentError("registration.full_weekend_pass.level", err.Error())
			case parseFieldFailurePassType:
				return nil, twirp.InvalidArgumentError("registration.pass_type", err.Error())
			case parseFieldFailureRole:
				return nil, twirp.InvalidArgumentError("registration.mix_and_match.role", err.Error())
			case parseFieldFailureStyle:
				return nil, twirp.InvalidArgumentError("registration.tshirt.style", err.Error())
			case parseFieldFailureHousing:
				return nil, twirp.InvalidArgumentError("registration.housing", err.Error())
			}
		}
		return nil, err
	}

	r, err := s.service.Update(ctx, auth, info)
	if err != nil {
		var noDiscountErr storage.ErrDiscountNotFound
		var noRegistrationErr storage.ErrNoRegistrationForID
		var hasSquarePayment registration.ErrHasSquarePayment
		var hasImmutableField registration.ErrHasImmutableField
		if errors.As(err, &noDiscountErr) {
			return nil, twirp.NewError(twirp.NotFound, noDiscountErr.Error()).WithMeta("Code", noDiscountErr.Code)
		} else if errors.As(err, &noRegistrationErr) {
			return nil, twirp.NewError(twirp.NotFound, noRegistrationErr.Error()).WithMeta("id", noRegistrationErr.ID)
		} else if errors.Is(err, registration.ErrRegistrationDisabled) {
			return nil, twirp.NewError(twirp.FailedPrecondition, "registration is disabled")
		} else if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		} else if errors.As(err, &hasSquarePayment) {
			return nil, twirp.InvalidArgumentError(hasSquarePayment.Field, hasSquarePayment.Error())
		} else if errors.As(err, &hasImmutableField) {
			return nil, twirp.InvalidArgumentError(hasImmutableField.Field, hasImmutableField.Error())
		}
		return nil, err
	}

	protoRegistration, err := toProtoc(r)
	if err != nil {
		return nil, fmt.Errorf("error transforming registration to protoc type: %w", err)
	}

	return &pb.RegistrationUpdateRes{
		Registration: protoRegistration,
	}, nil
}
