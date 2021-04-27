package forms

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/forms"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/golang/protobuf/ptypes"
	"github.com/twitchtv/twirp"
)

func (s *Server) SafetyReport(ctx context.Context, req *pb.SafetyReportReq) (*pb.SafetyReportRes, error) {
	occurredOn, err := ptypes.Timestamp(req.OccurredOn)
	if err != nil {
		return nil, twirp.InvalidArgumentError("occurred_on", fmt.Sprintf("error converting timestamp: %v", err))
	}
	if err := s.service.SafetyReport(ctx, occurredOn, req.Description, int(req.Severity), req.IssuesBefore, req.Resolution, req.Name, req.Email, req.PhoneNumber, req.RecaptchaResponse); err != nil {
		if errors.Is(err, forms.ErrRecaptchaFailed) {
			return nil, twirp.InvalidArgumentError("recaptcha_response", err.Error())
		}
	}

	return &pb.SafetyReportRes{}, nil
}
