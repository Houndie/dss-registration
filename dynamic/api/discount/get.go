package discount

import (
	"context"
	"errors"

	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

func (s *Server) Get(ctx context.Context, req *pb.DiscountGetReq) (*pb.DiscountGetRes, error) {
	bundle, err := s.service.Get(ctx, req.Code)
	if err != nil {
		e := storage.ErrDiscountNotFound{}
		if errors.As(err, &e) {
			return nil, twirp.NewError(twirp.NotFound, e.Error()).WithMeta("code", e.Code)
		}
		return nil, err
	}

	b, err := bundleToProto(req.Code, bundle)
	if err != nil {
		return nil, err
	}
	return &pb.DiscountGetRes{
		Bundle: b,
	}, nil
}
