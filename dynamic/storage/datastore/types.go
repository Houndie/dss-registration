package datastore

const (
	registrationKind = "Registration"

	fullWeekendPass = "Full Weekend Pass"
	danceOnlyPass   = "Dance Only Pass"
	noPass          = "No Pass"

	requiresHousing = "Requires Housing"
	providesHousing = "Provides Housing"
	noHousing       = "No Housing"
)

type registrationEntity struct {
	FirstName      string
	LastName       string
	StreetAddress  string
	City           string
	State          string
	ZipCode        string
	Email          string
	HomeScene      string
	IsStudent      bool
	SoloJazz       bool
	HousingRequest string
	RequireHousing struct {
		PetAllergies string
		Details      string
	}
	ProvideHousing struct {
		Pets     string
		Quantity int
		Details  string
	}
	WantsTShirt         bool
	TShirtStyle         string
	HasTeamCompetition  bool
	TeamCompetitionName string
	HasMixAndMatch      bool
	MixAndMatchRole     string
	WeekendPass         string
	FullWeekendPassInfo struct {
		Level int
		Tier  int
	}
	UserId string
}

const (
	orderKind = "Order"

	nonePayment                 = "none"
	manualPayment               = "manual"
	automaticPayment            = "automatic"
	completedButNoRecordPayment = "completed but no record"
)

type orderEntity struct {
	ReferenceId      string
	OrderId          string
	PaymentType      string
	AutomaticPayment string
}
