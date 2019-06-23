package square

type TenderCardDetailsStatus string

const (
	TenderCardDetailsStatusAuthorized TenderCardDetailsStatus = "AUTHORIZED"
	TenderCardDetailsStatusCaptured   TenderCardDetailsStatus = "CAPTURED"
	TenderCardDetailsStatusVoided     TenderCardDetailsStatus = "VOIDED"
	TenderCardDetailsStatusFailed     TenderCardDetailsStatus = "FAILED"
)

type TenderCardDetailsEntryMethod string

const (
	TenderCardDetailsEntryMethodSwiped      TenderCardDetailsEntryMethod = "SWIPED"
	TenderCardDetailsEntryMethodKeyed       TenderCardDetailsEntryMethod = "KEYED"
	TenderCardDetailsEntryMethodEmv         TenderCardDetailsEntryMethod = "EMV"
	TenderCardDetailsEntryMethodOnFile      TenderCardDetailsEntryMethod = "ON_FILE"
	TenderCardDetailsEntryMethodContactless TenderCardDetailsEntryMethod = "CONTACTLESS"
)

type TenderCardDetails struct {
	Status      TenderCardDetailsStatus      `json:"status"`
	Card        *Card                        `json:"card"`
	EntryMethod TenderCardDetailsEntryMethod `json:"entry_method"`
}

func (*TenderCardDetails) isTenderType() {}
