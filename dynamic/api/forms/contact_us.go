package forms

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/forms"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/twitchtv/twirp"
)

func (s *Server) ContactUs(ctx context.Context, req *pb.ContactUsReq) (*pb.ContactUsRes, error) {
	if err := s.service.ContactUs(ctx, req.Name, req.Email, req.Msg, req.RecaptchaResponse); err != nil {
		if err == forms.ErrRecaptchaFailed {
			return nil, twirp.InvalidArgumentError("recaptcha_response", err.Error())
		}
		return nil, err
	}

	return &pb.ContactUsRes{}, nil
}
