package registration

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/api"
	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/twitchtv/twirp"
)

type RegistrationService interface {
	Add(ctx context.Context, registration *registration.Info, redirectUrl, idempotencyKey, accessToken string) (string, error)
}

type Server struct {
	service RegistrationService
}

func (s *Server) Add(ctx context.Context, req *pb.RegistrationAddReq) (*pb.RegistrationAddRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		auth = ""
	}

	info, err := fromProtoc(req.Registration)
	if err != nil {
		var parseErr errParseField
		if errors.As(err, parseErr) {
			switch parseErr.field {
			case parseFieldFailureCreatedAt:
				return nil, twirp.InvalidArgumentError("registration.created_at", err.Error())
			case parseFieldFailureTier:
				return nil, twirp.InvalidArgumentError("registration.full_weekend_pass.tier", err.Error())
			case parseFieldFailureLevel:
				return nil, twirp.InvalidArgumentError("registration.full_weekend_pass.level", err.Error())
			case parseFieldFailurePassType:
				return nil, twirp.InvalidArgumentError("registration.pass_type", err.Error())
			case parseFieldFailureRole:
				return nil, twirp.InvalidArgumentError("registration.mix_and_match.role", err.Error())
			case parseFieldFailureStyle:
				return nil, twirp.InvalidArgumentError("registration.tshirt.style", err.Error())
			case parseFieldFailureHousing:
				return nil, twirp.InvalidArgumentError("registration.housing", err.Error())
			}
		}
		return nil, err
	}
	redirectURL, err := s.service.Add(ctx, info, req.RedirectUrl, req.IdempotencyKey, auth)
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationAddRes{
		RedirectUrl: redirectURL,
	}, nil

}

type parseFieldFailure int

const (
	parseFieldFailureCreatedAt parseFieldFailure = iota
	parseFieldFailureTier
	parseFieldFailureLevel
	parseFieldFailurePassType
	parseFieldFailureRole
	parseFieldFailureStyle
	parseFieldFailureHousing
)

func fromProtoc(r *pb.RegistrationInfo) (*registration.Info, error) {
	var passType registration.PassType
	switch pt := r.PassType.(type) {
	case *pb.RegistrationInfo_FullWeekendPass:
		var tier storage.WeekendPassTier
		switch pt.FullWeekendPass.Tier {
		case pb.FullWeekendPass_Tier1:
			tier = storage.Tier1
		case pb.FullWeekendPass_Tier2:
			tier = storage.Tier2
		case pb.FullWeekendPass_Tier3:
			tier = storage.Tier3
		case pb.FullWeekendPass_Tier4:
			tier = storage.Tier4
		case pb.FullWeekendPass_Tier5:
			tier = storage.Tier5
		default:
			return nil, errParseField{
				field: parseFieldFailureTier,
				msg:   fmt.Errorf("unexpected tier value %v", pb.FullWeekendPass_Tier_name[int32(pt.FullWeekendPass.Tier)]),
			}
		}

		var level storage.WeekendPassLevel
		switch pt.FullWeekendPass.Level {
		case pb.FullWeekendPass_Level1:
			level = storage.Level1
		case pb.FullWeekendPass_Level2:
			level = storage.Level2
		case pb.FullWeekendPass_Level3:
			level = storage.Level3
		default:
			return nil, errParseField{
				field: parseFieldFailureLevel,
				msg:   fmt.Errorf("unexpected level value %v", pb.FullWeekendPass_Level_name[int32(pt.FullWeekendPass.Level)]),
			}
		}

		passType = &registration.WeekendPass{
			Tier:  tier,
			Level: level,
			Paid:  pt.FullWeekendPass.Paid,
		}
	case *pb.RegistrationInfo_DanceOnlyPass:
		passType = &registration.DanceOnlyPass{
			Paid: pt.DanceOnlyPass.Paid,
		}
	case *pb.RegistrationInfo_NoPass:
		passType = &registration.NoPass{}
	default:
		return nil, errParseField{
			field: parseFieldFailurePassType,
			msg:   fmt.Errorf("unexpected pass type %T", r.PassType),
		}
	}

	var mixAndMatch *registration.MixAndMatch
	if r.MixAndMatch != nil {
		var role storage.MixAndMatchRole
		switch r.MixAndMatch.Role {
		case pb.MixAndMatch_Leader:
			role = storage.MixAndMatchRoleLeader
		case pb.MixAndMatch_Follower:
			role = storage.MixAndMatchRoleFollower
		default:
			return nil, errParseField{
				field: parseFieldFailureRole,
				msg:   fmt.Errorf("unexpected role value %v", pb.MixAndMatch_Role_name[int32(r.MixAndMatch.Role)]),
			}
		}
		mixAndMatch = &registration.MixAndMatch{
			Role: role,
			Paid: r.MixAndMatch.Paid,
		}
	}

	var soloJazz *registration.SoloJazz
	if r.SoloJazz != nil {
		soloJazz = &registration.SoloJazz{
			Paid: r.SoloJazz.Paid,
		}
	}

	var teamCompetition *registration.TeamCompetition
	if r.TeamCompetition != nil {
		teamCompetition = &registration.TeamCompetition{
			Name: r.TeamCompetition.Name,
			Paid: r.TeamCompetition.Paid,
		}
	}

	var tShirt *registration.TShirt
	if r.Tshirt != nil {
		var style storage.TShirtStyle
		switch r.Tshirt.Style {
		case pb.TShirt_UnisexS:
			style = storage.TShirtStyleUnisexS
		case pb.TShirt_UnisexM:
			style = storage.TShirtStyleUnisexM
		case pb.TShirt_UnisexL:
			style = storage.TShirtStyleUnisexL
		case pb.TShirt_UnisexXL:
			style = storage.TShirtStyleUnisexXL
		case pb.TShirt_Unisex2XL:
			style = storage.TShirtStyleUnisex2XL
		case pb.TShirt_Unisex3XL:
			style = storage.TShirtStyleUnisex3XL
		case pb.TShirt_BellaS:
			style = storage.TShirtStyleBellaS
		case pb.TShirt_BellaM:
			style = storage.TShirtStyleBellaM
		case pb.TShirt_BellaL:
			style = storage.TShirtStyleBellaL
		case pb.TShirt_BellaXL:
			style = storage.TShirtStyleBellaXL
		case pb.TShirt_Bella2XL:
			style = storage.TShirtStyleBella2XL
		default:
			return nil, errParseField{
				field: parseFieldFailureStyle,
				msg:   fmt.Errorf("unexpected style value %v", pb.TShirt_Style_name[int32(r.Tshirt.Style)]),
			}
		}
		tShirt = &registration.TShirt{
			Style: style,
			Paid:  r.Tshirt.Paid,
		}
	}

	var housing storage.Housing
	switch h := r.Housing.(type) {
	case *pb.RegistrationInfo_ProvideHousing:
		housing = &storage.ProvideHousing{
			Pets:     h.ProvideHousing.Pets,
			Quantity: int(h.ProvideHousing.Quantity),
			Details:  h.ProvideHousing.Details,
		}
	case *pb.RegistrationInfo_RequireHousing:
		housing = &storage.RequireHousing{
			PetAllergies: h.RequireHousing.PetAllergies,
			Details:      h.RequireHousing.Details,
		}
	case *pb.RegistrationInfo_NoHousing:
		housing = &storage.NoHousing{}
	default:
		return nil, errParseField{
			field: parseFieldFailureHousing,
			msg:   fmt.Errorf("unexpected housing type %T", r.Housing),
		}
	}

	createdAt, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		return nil, errParseField{
			field: parseFieldFailureCreatedAt,
			msg:   err,
		}
	}

	return &registration.Info{
		ID:              r.Id,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        soloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		Housing:         housing,
		DiscountCodes:   r.DiscountCodes,
		CreatedAt:       createdAt,
	}, nil
}

type errParseField struct {
	field parseFieldFailure
	msg   error
}

func (e errParseField) Error() string {
	return e.msg.Error()
}

func (e errParseField) Unwrap() error {
	return e.msg
}
