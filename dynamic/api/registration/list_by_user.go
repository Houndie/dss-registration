package registration

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) ListByUser(ctx context.Context, req *pb.RegistrationListByUserReq) (*pb.RegistrationListByUserRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}
	registrations, err := s.service.ListByUser(ctx, auth)
	if err != nil {
		if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		}
		return nil, err
	}
	results := make([]*pb.RegistrationInfo, len(registrations))
	for i, r := range registrations {
		results[i], err = toProtoc(r)
		if err != nil {
			return nil, fmt.Errorf("error converting registration to proto: %w", err)
		}
	}
	return &pb.RegistrationListByUserRes{
		Registrations: results,
	}, nil
}
