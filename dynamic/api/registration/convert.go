package registration

import (
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/registration"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

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

var (
	tierToProtoc = map[storage.WeekendPassTier]pb.FullWeekendPassTier{
		storage.Tier1: pb.FullWeekendPassTier_Tier1,
		storage.Tier2: pb.FullWeekendPassTier_Tier2,
		storage.Tier3: pb.FullWeekendPassTier_Tier3,
		storage.Tier4: pb.FullWeekendPassTier_Tier4,
		storage.Tier5: pb.FullWeekendPassTier_Tier5,
	}

	tierFromProtoc = map[pb.FullWeekendPassTier]storage.WeekendPassTier{
		pb.FullWeekendPassTier_Tier1: storage.Tier1,
		pb.FullWeekendPassTier_Tier2: storage.Tier2,
		pb.FullWeekendPassTier_Tier3: storage.Tier3,
		pb.FullWeekendPassTier_Tier4: storage.Tier4,
		pb.FullWeekendPassTier_Tier5: storage.Tier5,
	}

	levelToProtoc = map[storage.WeekendPassLevel]pb.FullWeekendPassLevel{
		storage.Level1: pb.FullWeekendPassLevel_Level1,
		storage.Level2: pb.FullWeekendPassLevel_Level2,
		storage.Level3: pb.FullWeekendPassLevel_Level3,
	}

	levelFromProtoc = map[pb.FullWeekendPassLevel]storage.WeekendPassLevel{
		pb.FullWeekendPassLevel_Level1: storage.Level1,
		pb.FullWeekendPassLevel_Level2: storage.Level2,
		pb.FullWeekendPassLevel_Level3: storage.Level3,
	}

	roleToProtoc = map[storage.MixAndMatchRole]pb.MixAndMatch_Role{
		storage.MixAndMatchRoleLeader:   pb.MixAndMatch_Leader,
		storage.MixAndMatchRoleFollower: pb.MixAndMatch_Follower,
	}

	roleFromProtoc = map[pb.MixAndMatch_Role]storage.MixAndMatchRole{
		pb.MixAndMatch_Leader:   storage.MixAndMatchRoleLeader,
		pb.MixAndMatch_Follower: storage.MixAndMatchRoleFollower,
	}

	styleToProtoc = map[storage.TShirtStyle]pb.TShirt_Style{
		storage.TShirtStyleUnisexS:   pb.TShirt_UnisexS,
		storage.TShirtStyleUnisexM:   pb.TShirt_UnisexM,
		storage.TShirtStyleUnisexL:   pb.TShirt_UnisexL,
		storage.TShirtStyleUnisexXL:  pb.TShirt_UnisexXL,
		storage.TShirtStyleUnisex2XL: pb.TShirt_Unisex2XL,
		storage.TShirtStyleUnisex3XL: pb.TShirt_Unisex3XL,
		storage.TShirtStyleBellaS:    pb.TShirt_BellaS,
		storage.TShirtStyleBellaM:    pb.TShirt_BellaM,
		storage.TShirtStyleBellaL:    pb.TShirt_BellaL,
		storage.TShirtStyleBellaXL:   pb.TShirt_BellaXL,
		storage.TShirtStyleBella2XL:  pb.TShirt_Bella2XL,
	}

	styleFromProtoc = map[pb.TShirt_Style]storage.TShirtStyle{
		pb.TShirt_UnisexS:   storage.TShirtStyleUnisexS,
		pb.TShirt_UnisexM:   storage.TShirtStyleUnisexM,
		pb.TShirt_UnisexL:   storage.TShirtStyleUnisexL,
		pb.TShirt_UnisexXL:  storage.TShirtStyleUnisexXL,
		pb.TShirt_Unisex2XL: storage.TShirtStyleUnisex2XL,
		pb.TShirt_Unisex3XL: storage.TShirtStyleUnisex3XL,
		pb.TShirt_BellaS:    storage.TShirtStyleBellaS,
		pb.TShirt_BellaM:    storage.TShirtStyleBellaM,
		pb.TShirt_BellaL:    storage.TShirtStyleBellaL,
		pb.TShirt_BellaXL:   storage.TShirtStyleBellaXL,
		pb.TShirt_Bella2XL:  storage.TShirtStyleBella2XL,
	}
)

func fromProtoc(r *pb.RegistrationInfo) (*registration.Info, error) {
	var passType registration.PassType
	switch pt := r.PassType.(type) {
	case *pb.RegistrationInfo_FullWeekendPass:
		tier, ok := tierFromProtoc[pt.FullWeekendPass.Tier]
		if !ok {
			return nil, errParseField{
				field: parseFieldFailureTier,
				msg:   fmt.Errorf("unexpected tier value %v", pb.FullWeekendPassTier_name[int32(pt.FullWeekendPass.Tier)]),
			}
		}

		level, ok := levelFromProtoc[pt.FullWeekendPass.Level]
		if !ok {
			return nil, errParseField{
				field: parseFieldFailureLevel,
				msg:   fmt.Errorf("unexpected level value %v", pb.FullWeekendPassLevel_name[int32(pt.FullWeekendPass.Level)]),
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
		role, ok := roleFromProtoc[r.MixAndMatch.Role]
		if !ok {
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
		style, ok := styleFromProtoc[r.Tshirt.Style]
		if !ok {
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

func toProtoc(r *registration.Info) (*pb.RegistrationInfo, error) {
	result := &pb.RegistrationInfo{
		Id:            r.ID,
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		StreetAddress: r.StreetAddress,
		City:          r.City,
		State:         r.State,
		ZipCode:       r.ZipCode,
		Email:         r.Email,
		HomeScene:     r.HomeScene,
		IsStudent:     r.IsStudent,
		DiscountCodes: r.DiscountCodes,
		CreatedAt:     r.CreatedAt.Format(time.RFC3339),
	}

	switch pt := r.PassType.(type) {
	case *registration.WeekendPass:
		tier, ok := tierToProtoc[pt.Tier]
		if !ok {
			return nil, fmt.Errorf("unknown tier %s", tier)
		}
		level, ok := levelToProtoc[pt.Level]
		if !ok {
			return nil, fmt.Errorf("unknown level %s", level)
		}
		result.PassType = &pb.RegistrationInfo_FullWeekendPass{
			FullWeekendPass: &pb.FullWeekendPass{
				Tier:  tier,
				Level: level,
				Paid:  pt.Paid,
			},
		}
	case *registration.DanceOnlyPass:
		result.PassType = &pb.RegistrationInfo_DanceOnlyPass{
			DanceOnlyPass: &pb.DanceOnlyPass{
				Paid: pt.Paid,
			},
		}
	case *registration.NoPass:
		result.PassType = &pb.RegistrationInfo_NoPass{
			NoPass: &pb.NoPass{},
		}
	default:
		return nil, fmt.Errorf("unknown pass type %T", pt)
	}

	if r.MixAndMatch != nil {
		role, ok := roleToProtoc[r.MixAndMatch.Role]
		if !ok {
			return nil, fmt.Errorf("unknown role %s", role)
		}
		result.MixAndMatch = &pb.MixAndMatch{
			Role: role,
			Paid: r.MixAndMatch.Paid,
		}
	}

	if r.SoloJazz != nil {
		result.SoloJazz = &pb.SoloJazz{
			Paid: r.SoloJazz.Paid,
		}
	}

	if r.TeamCompetition != nil {
		result.TeamCompetition = &pb.TeamCompetition{
			Name: r.TeamCompetition.Name,
			Paid: r.TeamCompetition.Paid,
		}
	}

	switch h := r.Housing.(type) {
	case *storage.ProvideHousing:
		result.Housing = &pb.RegistrationInfo_ProvideHousing{
			ProvideHousing: &pb.ProvideHousing{
				Pets:     h.Pets,
				Quantity: int64(h.Quantity),
				Details:  h.Details,
			},
		}
	case *storage.RequireHousing:
		result.Housing = &pb.RegistrationInfo_RequireHousing{
			RequireHousing: &pb.RequireHousing{
				PetAllergies: h.PetAllergies,
				Details:      h.Details,
			},
		}
	case *storage.NoHousing:
		result.Housing = &pb.RegistrationInfo_NoHousing{
			NoHousing: &pb.NoHousing{},
		}
	default:
		return nil, fmt.Errorf("unknown housing type %T", h)
	}

	return result, nil
}
