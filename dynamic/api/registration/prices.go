package registration

import (
	"context"
	"fmt"

	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
)

func (s *Server) Prices(ctx context.Context, req *pb.RegistrationPricesReq) (*pb.RegistrationPricesRes, error) {
	prices, err := s.service.Populate(ctx)
	if err != nil {
		return nil, err
	}

	tier, ok := tierToProtoc[prices.WeekendPassTier]
	if !ok {
		return nil, fmt.Errorf("unknown weekend pass tier %v", tier)
	}
	studentDiscount, err := discountToProtoc(prices.StudentDiscount)
	if err != nil {
		return nil, fmt.Errorf("error converting student discount to protoc discount: %w", err)
	}
	return &pb.RegistrationPricesRes{
		WeekendPassCost:     int64(prices.WeekendPassCost),
		WeekendPassTier:     tier,
		DancePassCost:       int64(prices.DancePassCost),
		MixAndMatchCost:     int64(prices.MixAndMatchCost),
		SoloJazzCost:        int64(prices.SoloJazzCost),
		TeamCompetitionCost: int64(prices.TeamCompCost),
		TshirtCost:          int64(prices.TShirtCost),
		StudentDiscount:     studentDiscount,
	}, nil
}
