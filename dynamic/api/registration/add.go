package registration

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

func (s *Server) Add(ctx context.Context, req *pb.RegistrationAddReq) (*pb.RegistrationAddRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		auth = ""
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
	redirectURL, err := s.service.Add(ctx, info, req.RedirectUrl, req.IdempotencyKey, auth)
	if err != nil {
		var noDiscountErr storage.ErrDiscountNotFound
		var outOfStockErr registration.ErrOutOfStock
		if errors.As(err, &noDiscountErr) {
			return nil, twirp.InvalidArgumentError("registration.discounts", noDiscountErr.Error()).WithMeta("Code", noDiscountErr.Code)
		} else if errors.Is(err, registration.ErrRegistrationDisabled) {
			return nil, twirp.NewError(twirp.FailedPrecondition, "registration is disabled")
		} else if errors.As(err, &outOfStockErr) {
			return nil, twirp.NewError(twirp.FailedPrecondition, outOfStockErr.Error()).WithMeta("next_tier", string(outOfStockErr.NextTier)).WithMeta("next_cost", string(outOfStockErr.NextCost))
		}
		return nil, err
	}
	return &pb.RegistrationAddRes{
		RedirectUrl: redirectURL,
	}, nil

}
