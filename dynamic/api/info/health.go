package info

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/info"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
)

func (s *Server) Health(ctx context.Context, req *pb.InfoHealthReq) (*pb.InfoHealthRes, error) {
	res := &pb.InfoHealthRes{}
	switch s.service.Health(ctx) {
	case info.Healthy:
		res.Healthiness = pb.InfoHealthRes_Healthy
	case info.Unhealthy:
		res.Healthiness = pb.InfoHealthRes_Unhealthy
	}

	return res, nil
}
