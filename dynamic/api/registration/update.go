package registration

import (
	"context"
	"errors"

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
		if errors.As(err, parseErr) {
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

	redirectURL, err := s.service.Update(ctx, auth, req.IdempotencyKey, info, req.RedirectUrl)
	if err != nil {
		var noDiscountErr storage.ErrDiscountNotFound
		var outOfStockErr registration.ErrOutOfStock
		if errors.As(err, &noDiscountErr) {
			return nil, twirp.InvalidArgumentError("registration.discounts", noDiscountErr.Error()).WithMeta("Code", noDiscountErr.Code)
		} else if errors.Is(err, registration.ErrRegistrationDisabled) {
			return nil, twirp.NewError(twirp.FailedPrecondition, "registration is disabled")
		} else if errors.As(err, &outOfStockErr) {
			return nil, twirp.NewError(twirp.FailedPrecondition, outOfStockErr.Error()).WithMeta("next_tier", string(outOfStockErr.NextTier)).WithMeta("next_cost", string(outOfStockErr.NextCost))
		} else if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		}
		return nil, err
	}
	return &pb.RegistrationUpdateRes{
		RedirectUrl: redirectURL,
	}, nil
}
