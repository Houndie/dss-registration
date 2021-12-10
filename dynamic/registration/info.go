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
	Level                storage.WeekendPassLevel
	Tier                 storage.WeekendPassTier
	SquarePaid           bool
	AdminPaymentOverride bool
}
type DanceOnlyPass struct {
	SquarePaid           bool
	AdminPaymentOverride bool
}
type NoPass struct{}

func (*WeekendPass) isPassType()   {}
func (*DanceOnlyPass) isPassType() {}
func (*NoPass) isPassType()        {}

func toStoragePassType(passType PassType) storage.PassType {
	switch p := passType.(type) {
	case *WeekendPass:
		return &storage.WeekendPass{
			Level:        p.Level,
			Tier:         p.Tier,
			ManuallyPaid: p.AdminPaymentOverride,
		}
	case *DanceOnlyPass:
		return &storage.DanceOnlyPass{
			ManuallyPaid: p.AdminPaymentOverride,
		}
	}
	return &storage.NoPass{}
}

func fromStoragePassType(passType storage.PassType, paidWeekend, paidDance bool) PassType {
	switch p := passType.(type) {
	case *storage.WeekendPass:
		return &WeekendPass{
			Level:                p.Level,
			Tier:                 p.Tier,
			SquarePaid:           paidWeekend,
			AdminPaymentOverride: p.ManuallyPaid,
		}
	case *storage.DanceOnlyPass:

		return &DanceOnlyPass{
			SquarePaid:           paidDance,
			AdminPaymentOverride: p.ManuallyPaid,
		}
	}
	return &NoPass{}
}

type MixAndMatch struct {
	Role                 storage.MixAndMatchRole
	SquarePaid           bool
	AdminPaymentOverride bool
}

func toStorageMixAndMatch(m *MixAndMatch) *storage.MixAndMatch {
	if m == nil {
		return nil
	}
	return &storage.MixAndMatch{
		Role:         m.Role,
		ManuallyPaid: m.AdminPaymentOverride,
	}
}

func fromStorageMixAndMatch(m *storage.MixAndMatch, squarePaid bool) *MixAndMatch {
	if m == nil {
		return nil
	}

	return &MixAndMatch{
		Role:                 m.Role,
		SquarePaid:           squarePaid,
		AdminPaymentOverride: m.ManuallyPaid,
	}
}

type SoloJazz struct {
	SquarePaid           bool
	AdminPaymentOverride bool
}

func toStorageSoloJazz(s *SoloJazz) *storage.SoloJazz {
	if s == nil {
		return nil
	}
	return &storage.SoloJazz{
		ManuallyPaid: s.AdminPaymentOverride,
	}
}

func fromStorageSoloJazz(s *storage.SoloJazz, squarePaid bool) *SoloJazz {
	if s == nil {
		return nil
	}

	return &SoloJazz{
		SquarePaid:           squarePaid,
		AdminPaymentOverride: s.ManuallyPaid,
	}
}

type TeamCompetition struct {
	Name                 string
	SquarePaid           bool
	AdminPaymentOverride bool
}

func toStorageTeamCompetition(t *TeamCompetition) *storage.TeamCompetition {
	if t == nil {
		return nil
	}
	return &storage.TeamCompetition{
		Name:         t.Name,
		ManuallyPaid: t.AdminPaymentOverride,
	}
}

func fromStorageTeamCompetition(t *storage.TeamCompetition, squarePaid bool) *TeamCompetition {
	if t == nil {
		return nil
	}

	return &TeamCompetition{
		Name:                 t.Name,
		SquarePaid:           squarePaid,
		AdminPaymentOverride: t.ManuallyPaid,
	}
}

type TShirt struct {
	Style                storage.TShirtStyle
	SquarePaid           bool
	AdminPaymentOverride bool
}

func toStorageTShirt(t *TShirt) *storage.TShirt {
	if t == nil {
		return nil
	}
	return &storage.TShirt{
		Style:        t.Style,
		ManuallyPaid: t.AdminPaymentOverride,
	}
}

func fromStorageTShirt(t *storage.TShirt, squarePaid bool) *TShirt {
	if t == nil {
		return nil
	}

	return &TShirt{
		Style:                t.Style,
		SquarePaid:           squarePaid,
		AdminPaymentOverride: t.ManuallyPaid,
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
		Enabled:         r.Enabled,
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
		Enabled:         r.Enabled,
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
	Enabled         bool
}
