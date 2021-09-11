package registration

import (
	"time"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

type PassType interface {
	isPassType()
}
type WeekendPass struct {
	Level storage.WeekendPassLevel
	Tier  storage.WeekendPassTier
	Paid  bool
}
type DanceOnlyPass struct {
	Paid bool
}
type NoPass struct{}

func (*WeekendPass) isPassType()   {}
func (*DanceOnlyPass) isPassType() {}
func (*NoPass) isPassType()        {}

func toStoragePassType(passType PassType) storage.PassType {
	switch p := passType.(type) {
	case *WeekendPass:
		return &storage.WeekendPass{
			Level: p.Level,
			Tier:  p.Tier,
		}
	case *DanceOnlyPass:
		return &storage.DanceOnlyPass{}
	}
	return &storage.NoPass{}
}

func fromStoragePassType(passType storage.PassType, paidWeekend, paidDance bool) PassType {
	switch p := passType.(type) {
	case *storage.WeekendPass:
		return &WeekendPass{
			Level: p.Level,
			Tier:  p.Tier,
			Paid:  paidWeekend,
		}
	case *storage.DanceOnlyPass:
		return &DanceOnlyPass{
			Paid: paidDance,
		}
	}
	return &NoPass{}
}

type MixAndMatch struct {
	Role storage.MixAndMatchRole
	Paid bool
}

func toStorageMixAndMatch(m *MixAndMatch) *storage.MixAndMatch {
	if m == nil {
		return nil
	}
	return &storage.MixAndMatch{
		Role: m.Role,
	}
}

func fromStorageMixAndMatch(m *storage.MixAndMatch, paid bool) *MixAndMatch {
	if m == nil {
		return nil
	}
	return &MixAndMatch{
		Role: m.Role,
		Paid: paid,
	}
}

type SoloJazz struct {
	Paid bool
}

func toStorageSoloJazz(s *SoloJazz) bool {
	return s != nil
}

func fromStorageSoloJazz(s bool, paid bool) *SoloJazz {
	if !s {
		return nil
	}
	return &SoloJazz{
		Paid: paid,
	}
}

type TeamCompetition struct {
	Name string
	Paid bool
}

func toStorageTeamCompetition(t *TeamCompetition) *storage.TeamCompetition {
	if t == nil {
		return nil
	}
	return &storage.TeamCompetition{
		Name: t.Name,
	}
}

func fromStorageTeamCompetition(t *storage.TeamCompetition, paid bool) *TeamCompetition {
	if t == nil {
		return nil
	}
	return &TeamCompetition{
		Name: t.Name,
		Paid: paid,
	}
}

type TShirt struct {
	Style storage.TShirtStyle
	Paid  bool
}

func toStorageTShirt(t *TShirt) *storage.TShirt {
	if t == nil {
		return nil
	}
	return &storage.TShirt{
		Style: t.Style,
	}
}

func fromStorageTShirt(t *storage.TShirt, paid bool) *TShirt {
	if t == nil {
		return nil
	}
	return &TShirt{
		Style: t.Style,
		Paid:  paid,
	}
}

func toStorageRegistration(r *Info) *storage.Registration {
	return &storage.Registration{
		ID:              r.ID,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        toStoragePassType(r.PassType),
		MixAndMatch:     toStorageMixAndMatch(r.MixAndMatch),
		SoloJazz:        toStorageSoloJazz(r.SoloJazz),
		TeamCompetition: toStorageTeamCompetition(r.TeamCompetition),
		TShirt:          toStorageTShirt(r.TShirt),
		Housing:         r.Housing,
		DiscountCodes:   r.DiscountCodes,
		CreatedAt:       r.CreatedAt,
	}
}

func fromStorageRegistration(r *storage.Registration, pd *common.PaymentData) *Info {
	return &Info{
		ID:              r.ID,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        fromStoragePassType(r.PassType, pd.WeekendPassPaid, pd.DanceOnlyPaid),
		MixAndMatch:     fromStorageMixAndMatch(r.MixAndMatch, pd.MixAndMatchPaid),
		SoloJazz:        fromStorageSoloJazz(r.SoloJazz, pd.SoloJazzPaid),
		TeamCompetition: fromStorageTeamCompetition(r.TeamCompetition, pd.TeamCompetitionPaid),
		TShirt:          fromStorageTShirt(r.TShirt, pd.TShirtPaid),
		Housing:         r.Housing,
		DiscountCodes:   r.DiscountCodes,
		CreatedAt:       r.CreatedAt,
	}
}

type Info struct {
	ID              string
	FirstName       string
	LastName        string
	StreetAddress   string
	City            string
	State           string
	ZipCode         string
	Email           string
	HomeScene       string
	IsStudent       bool
	PassType        PassType
	MixAndMatch     *MixAndMatch
	SoloJazz        *SoloJazz
	TeamCompetition *TeamCompetition
	TShirt          *TShirt
	Housing         storage.Housing
	DiscountCodes   []string
	CreatedAt       time.Time
}
