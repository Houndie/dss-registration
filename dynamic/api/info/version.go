package info

import (
	"context"

	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
)

func (s *Server) Version(ctx context.Context, req *pb.InfoVersionReq) (*pb.InfoVersionRes, error) {
	return &pb.InfoVersionRes{
		Version: s.service.Version(),
	}, nil
}
