package discount

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) List(ctx context.Context, req *pb.DiscountListReq) (*pb.DiscountListRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, authorizer.Unauthenticated.Error())
	}
	bundles, err := s.service.List(ctx, auth)
	if err != nil {
		if errors.Is(err, common.ErrUnauthorized) {
			return nil, twirp.NewError(twirp.PermissionDenied, err.Error())
		} else if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, authorizer.Unauthenticated.Error())
		}
		return nil, err
	}

	protoBundles := make([]*pb.DiscountBundle, len(bundles))
	for i, bundle := range bundles {
		protoBundles[i], err = bundleToProto(bundle)
		if err != nil {
			return nil, err
		}
	}

	return &pb.DiscountListRes{
		Bundles: protoBundles,
	}, nil
}
