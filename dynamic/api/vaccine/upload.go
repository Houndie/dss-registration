package vaccine

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/api"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/vaccine"
	"github.com/twitchtv/twirp"
)

func (s *Server) Upload(ctx context.Context, req *pb.VaccineUploadReq) (*pb.VaccineUploadRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	u, err := s.service.Upload(ctx, auth, req.Filesize, req.Id)
	if err != nil {
		var fileTooBig vaccine.ErrFileTooBig
		if errors.As(err, &fileTooBig) {
			return nil, twirp.InvalidArgumentError("filesize", err.Error())
		}
		return nil, err
	}

	return &pb.VaccineUploadRes{
		Url: u,
	}, nil
}
