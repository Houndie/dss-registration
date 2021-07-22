package registration

import (
	"context"
	"fmt"

	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
)

func (s *Server) Prices(ctx context.Context, req *pb.RegistrationPricesReq) (*pb.RegistrationPricesRes, error) {
	tier, err := s.service.Populate(ctx)
	if err != nil {
		return nil, err
	}

	protoTier, ok := tierToProtoc[tier]
	if !ok {
		return nil, fmt.Errorf("unknown weekend pass tier %v", tier)
	}
	return &pb.RegistrationPricesRes{
		WeekendPassTier: protoTier,
	}, nil
}
