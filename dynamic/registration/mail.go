package registration

import (
	"errors"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

type mailPassTypeEnum string

const (
	mailWeekendPass   mailPassTypeEnum = "FULL_WEEKEND"
	mailDanceOnlyPass mailPassTypeEnum = "DANCE_ONLY"
	mailNoPass        mailPassTypeEnum = "NO_PASS"
)

type mailPassType struct {
	Type        mailPassTypeEnum     `json:"type,omitempty"`
	WeekendPass *mailWeekendPassData `json:"weekend_pass,omitempty"`
}

type mailWeekendPassData struct {
	Level int `json:"level,omitempty"`
}

type mailMixAndMatchRole string

const (
	mailLeader   mailMixAndMatchRole = "Leader"
	mailFollower mailMixAndMatchRole = "Follower"
)

type mailMixAndMatch struct {
	Purchased bool                `json:"purchased,omitempty"`
	Role      mailMixAndMatchRole `json:"role,omitempty"`
}

type mailSoloJazz struct {
	Purchased bool `json:"purchased,omitempty"`
}

type mailTeamComp struct {
	Purchased bool   `json:"purchased,omitempty"`
	Name      string `json:"name,omitempty"`
}

type mailTShirtStyle string

const (
	mailUnisexS   mailTShirtStyle = "Unisex S"
	mailUnisexM   mailTShirtStyle = "Unisex M"
	mailUnisexL   mailTShirtStyle = "Unisex L"
	mailUnisexXL  mailTShirtStyle = "Unisex XL"
	mailUnisex2XL mailTShirtStyle = "Unisex 2XL"
	mailUnisex3XL mailTShirtStyle = "Unisex 3XL"
	mailBellaS    mailTShirtStyle = "Bella S"
	mailBellaM    mailTShirtStyle = "Bella M"
	mailBellaL    mailTShirtStyle = "Bella L"
	mailBellaXL   mailTShirtStyle = "Bella XL"
	mailBella2XL  mailTShirtStyle = "Bella 2XL"
)

type mailTShirt struct {
	Purchased bool            `json:"purchased,omitempty"`
	Style     mailTShirtStyle `json:"style,omitempty"`
}

type mailHousingEnum string

const (
	mailProvideHousing mailHousingEnum = "PROVIDE"
	mailRequireHousing mailHousingEnum = "REQUIRE"
	mailNoHousing      mailHousingEnum = "NONE"
)

type mailHousing struct {
	Type    mailHousingEnum         `json:"type,omitempty"`
	Provide *mailProvideHousingData `json:"provide,omitempty"`
	Require *mailRequireHousingData `json:"require,omitempty"`
}

type mailProvideHousingData struct {
	Pets     string `json:"pets,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Details  string `json:"details,omitempty"`
}
type mailRequireHousingData struct {
	PetAllergies string `json:"pet_allergies"`
	Details      string `json:"details"`
}

type mailParams struct {
	FirstName       string           `json:"first_name"`
	LastName        string           `json:"last_name"`
	StreetAddress   string           `json:"street_address"`
	City            string           `json:"city"`
	State           string           `json:"state"`
	ZipCode         string           `json:"zip_code"`
	HomeScene       string           `json:"home_scene"`
	IsStudent       bool             `json:"is_student"`
	PassType        *mailPassType    `json:"pass_type,omitempty"`
	MixAndMatch     *mailMixAndMatch `json:"mix_and_match,omitempty"`
	SoloJazz        *mailSoloJazz    `json:"solo_jazz,omitempty"`
	TeamCompetition *mailTeamComp    `json:"team_competition,omitempty"`
	TShirt          *mailTShirt      `json:"t_shirt,omitempty"`
	Housing         *mailHousing     `json:"housing,omitempty"`
}

func toMailPassType(p PassType) (*mailPassType, error) {
	switch pt := p.(type) {
	case *WeekendPass:
		return &mailPassType{
			Type: mailWeekendPass,
			WeekendPass: &mailWeekendPassData{
				Level: int(pt.Level),
			},
		}, nil
	case *DanceOnlyPass:
		return &mailPassType{
			Type: mailDanceOnlyPass,
		}, nil
	case *NoPass:
		return &mailPassType{
			Type: mailNoPass,
		}, nil
	}
	return nil, errors.New("error converting weekend pass, unknown type")
}

func toMailMixAndMatch(m *MixAndMatch) *mailMixAndMatch {
	if m == nil {
		return &mailMixAndMatch{
			Purchased: false,
		}
	}
	var role mailMixAndMatchRole
	switch m.Role {
	case storage.MixAndMatchRoleLeader:
		role = mailLeader
	case storage.MixAndMatchRoleFollower:
		role = mailFollower
	}
	return &mailMixAndMatch{
		Purchased: true,
		Role:      role,
	}
}

func toMailSoloJazz(s *SoloJazz) *mailSoloJazz {
	if s == nil {
		return &mailSoloJazz{
			Purchased: false,
		}
	}
	return &mailSoloJazz{
		Purchased: true,
	}
}

func toMailTeamCompetition(t *TeamCompetition) *mailTeamComp {
	if t == nil {
		return &mailTeamComp{
			Purchased: false,
		}
	}
	return &mailTeamComp{
		Purchased: true,
		Name:      t.Name,
	}
}

func toMailTShirt(t *TShirt) *mailTShirt {
	if t == nil {
		return &mailTShirt{
			Purchased: false,
		}
	}
	var style mailTShirtStyle
	switch t.Style {
	case storage.TShirtStyleUnisexS:
		style = mailUnisexS
	case storage.TShirtStyleUnisexM:
		style = mailUnisexM
	case storage.TShirtStyleUnisexL:
		style = mailUnisexL
	case storage.TShirtStyleUnisexXL:
		style = mailUnisexXL
	case storage.TShirtStyleUnisex2XL:
		style = mailUnisex2XL
	case storage.TShirtStyleUnisex3XL:
		style = mailUnisex3XL
	case storage.TShirtStyleBellaS:
		style = mailBellaS
	case storage.TShirtStyleBellaM:
		style = mailBellaM
	case storage.TShirtStyleBellaL:
		style = mailBellaL
	case storage.TShirtStyleBellaXL:
		style = mailBellaXL
	case storage.TShirtStyleBella2XL:
		style = mailBella2XL
	}
	return &mailTShirt{
		Purchased: true,
		Style:     style,
	}
}

func toMailHousing(h storage.Housing) (*mailHousing, error) {
	switch ht := h.(type) {
	case *storage.NoHousing:
		return &mailHousing{
			Type: mailNoHousing,
		}, nil
	case *storage.ProvideHousing:
		return &mailHousing{
			Type: mailProvideHousing,
			Provide: &mailProvideHousingData{
				Pets:     ht.Pets,
				Quantity: ht.Quantity,
				Details:  ht.Details,
			},
		}, nil
	case *storage.RequireHousing:
		return &mailHousing{
			Type: mailRequireHousing,
			Require: &mailRequireHousingData{
				PetAllergies: ht.PetAllergies,
				Details:      ht.Details,
			},
		}, nil
	}
	return nil, errors.New("error converting to mail housing, unknown type")
}

func toMailParams(r *Info) (*mailParams, error) {
	passType, err := toMailPassType(r.PassType)
	if err != nil {
		return nil, err
	}

	housing, err := toMailHousing(r.Housing)
	if err != nil {
		return nil, err
	}

	return &mailParams{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        passType,
		MixAndMatch:     toMailMixAndMatch(r.MixAndMatch),
		SoloJazz:        toMailSoloJazz(r.SoloJazz),
		TeamCompetition: toMailTeamCompetition(r.TeamCompetition),
		TShirt:          toMailTShirt(r.TShirt),
		Housing:         housing,
	}, nil

}
