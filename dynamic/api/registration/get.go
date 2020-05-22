package registration

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/authorizer"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

func (s *Server) Get(ctx context.Context, req *pb.RegistrationGetReq) (*pb.RegistrationGetRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	r, err := s.service.Get(ctx, auth, req.Id)
	if err != nil {
		return nil, err
	}
	protoRegistration, err := toProtoc(r)
	if err != nil {
		var noRegistrationErr storage.ErrNoRegistrationForID
		if errors.Is(err, authorizer.Unauthenticated) {
			return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
		} else if errors.As(err, &noRegistrationErr) {
			return nil, twirp.InvalidArgumentError("id", noRegistrationErr.Error()).WithMeta("id", req.Id)
		}
		return nil, fmt.Errorf("error transforming registration to protoc type: %w", err)
	}

	return &pb.RegistrationGetRes{
		Registration: protoRegistration,
	}, nil
}
