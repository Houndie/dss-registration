package vaccine

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/api"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) Approve(ctx context.Context, req *pb.VaccineApproveReq) (*pb.VaccineApproveRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	err := s.service.Approve(ctx, auth, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.VaccineApproveRes{}, nil
}
