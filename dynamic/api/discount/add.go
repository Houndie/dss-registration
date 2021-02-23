package discount

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/discount"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) Add(ctx context.Context, req *pb.DiscountAddReq) (*pb.DiscountAddRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	d, err := bundleFromProto(req.Bundle)
	if err != nil {
		e1 := ErrUnknownDiscountType{}
		if errors.As(err, &e1) {
			return nil, twirp.InvalidArgumentError("bundle.discounts.amount", err.Error())
		}
		e2 := ErrUnknownPurchaseItem{}
		if errors.As(err, &e2) {
			return nil, twirp.InvalidArgumentError("bundle.discounts.appliedTo", err.Error())
		}
		return nil, err
	}
	if d.Code == "" {
		return nil, twirp.InvalidArgumentError("bundle.code", "value should be non-empty")
	}
	for _, discount := range d.Discounts {
		if discount.Name == "" {
			return nil, twirp.InvalidArgumentError("bundle.discounts.name", "value should be non-empty")
		}
		if discount.Amount != nil {
			return nil, twirp.InvalidArgumentError("bundle.discounts.amount", "value should not be provided")
		}
	}
	err = s.service.Add(ctx, auth, d)
	if err != nil {
		if errors.Is(err, discount.ErrUnauthorized) {
			return nil, twirp.NewError(twirp.PermissionDenied, err.Error())
		}
		return nil, err
	}
	return &pb.DiscountAddRes{}, nil
}
