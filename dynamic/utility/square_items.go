package utility

import "github.com/Houndie/dss-registration/dynamic/storage"

const (
	MixAndMatchItem = "Mix and Match"
	TeamCompItem    = "Team Competition"
	SoloJazzItem    = "Solo"
	DancePassItem   = "Dance Only"
	WeekendPassItem = "Full Weekend Pass"
	TShirtItem      = "2020 T-Shirt"

	StudentDiscountItem  = "Student Discount"
	DancePassPresaleName = "Presale"
)

var WeekendPassName = map[storage.WeekendPassTier]string{
	storage.Tier1: "Tier 1",
	storage.Tier2: "Tier 2",
	storage.Tier3: "Tier 3",
	storage.Tier4: "Tier 4",
	storage.Tier5: "Tier 5",
}

const SmackdownEmail = "info@daytonswingsmackdown.com"
