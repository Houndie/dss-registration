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

func (s *Server) Pay(ctx context.Context, req *pb.RegistrationPayReq) (*pb.RegistrationPayRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	checkoutURL, err := s.service.Pay(ctx, req.Id, req.RedirectUrl, req.IdempotencyKey, auth)
	if err != nil {
		var outOfStockErr registration.ErrOutOfStock
		var noRegistrationErr storage.ErrNoRegistrationForID
		if errors.As(err, &noRegistrationErr) {
			return nil, twirp.NewError(twirp.NotFound, noRegistrationErr.Error()).WithMeta("id", noRegistrationErr.ID)
		} else if errors.As(err, &outOfStockErr) {
			return nil, twirp.NewError(twirp.FailedPrecondition, outOfStockErr.Error()).WithMeta("next_tier", fmt.Sprintf("%v", outOfStockErr.NextTier)).WithMeta("next_cost", fmt.Sprintf("%v", outOfStockErr.NextCost))
		} else if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		}
		return nil, err
	}

	return &pb.RegistrationPayRes{
		CheckoutUrl: checkoutURL,
	}, nil
}
