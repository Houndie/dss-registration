package registration

import (
	"context"
	"errors"
	"time"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetSummary(ctx context.Context, req *pb.RegistrationGetSummaryReq) (*pb.RegistrationGetSummaryRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}
	summaries, err := s.service.SummaryByUser(ctx, auth)
	if err != nil {
		if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		}
		return nil, err
	}
	results := make([]*pb.RegistrationSummary, len(summaries))
	for i, summary := range summaries {
		results[i] = &pb.RegistrationSummary{
			Id:        summary.ID,
			FirstName: summary.FirstName,
			LastName:  summary.LastName,
			Email:     summary.Email,
			CreatedAt: summary.CreatedAt.Format(time.RFC3339),
			Paid:      summary.Paid,
		}
	}
	return &pb.RegistrationGetSummaryRes{
		Summaries: results,
	}, nil
}
