package vaccine

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/api"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) Reject(ctx context.Context, req *pb.VaccineRejectReq) (*pb.VaccineRejectRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	err := s.service.Reject(ctx, auth, req.Id, req.Reason)
	if err != nil {
		return nil, err
	}

	return &pb.VaccineRejectRes{}, nil
}
