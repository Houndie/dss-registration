package registration

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) UploadVaxImage(ctx context.Context, req *pb.RegistrationUploadVaxImageReq) (*pb.RegistrationUploadVaxImageRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	u, err := s.service.UploadVaxImage(ctx, auth, req.Filesize, req.Id)
	if err != nil {
		var fileTooBig registration.ErrFileTooBig
		if errors.As(err, &fileTooBig) {
			return nil, twirp.InvalidArgumentError("filesize", err.Error())
		}
		return nil, err
	}

	return &pb.RegistrationUploadVaxImageRes{
		Url: u,
	}, nil
}
