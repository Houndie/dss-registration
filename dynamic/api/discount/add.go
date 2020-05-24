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
		e := ErrUnknownDiscountType{}
		if errors.As(err, &e) {
			return nil, twirp.InvalidArgumentError("bundle.discounts.amount", err.Error())
		}
		return nil, err
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
