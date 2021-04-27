package discount

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

func (s *Server) Delete(ctx context.Context, req *pb.DiscountDeleteReq) (*pb.DiscountDeleteRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	err := s.service.Delete(ctx, auth, req.Code)
	if err != nil {
		if errors.Is(err, common.ErrUnauthorized) {
			return nil, twirp.NewError(twirp.PermissionDenied, err.Error())
		} else if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, authorizer.Unauthenticated.Error())
		}
		e := storage.ErrDiscountNotFound{}
		if errors.As(err, &e) {
			return nil, twirp.NewError(twirp.NotFound, e.Error()).WithMeta("code", e.Code)
		}
		return nil, err
	}
	return &pb.DiscountDeleteRes{}, nil
}
