package square

import (
	"encoding/json"
	"time"

	"github.com/peterhellberg/duration"
	"github.com/pkg/errors"
)

type orderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient `json:"recipient"`
	ExpiresAt            time.Time                  `json:"expires_at"`
	AutoCompleteDuration string                     `json:"auto_complete_duration"`
	PickupAt             time.Time                  `json:"pickup_at"`
	PickupWindowDuration string                     `json:"pickup_window_duration"`
	PrepTimeDuration     string                     `json:"prep_time_duration"`
	Note                 string                     `json:"note"`
	PlacedAt             time.Time                  `json:"placed_at"`
	AcceptedAt           time.Time                  `json:"accepted_at"`
	RejectedAt           time.Time                  `json:"rejected_at"`
	ReadyAt              time.Time                  `json:"ready_at"`
	ExpiredAt            time.Time                  `json:"expired_at"`
	PickedUpAt           time.Time                  `json:"picked_up_at"`
	CanceledAt           time.Time                  `json:"canceled_at"`
	CancelReason         string                     `json:"cancel_reason"`
}

type OrderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient `json:"recipient"`
	ExpiresAt            time.Time                  `json:"expires_at"`
	AutoCompleteDuration time.Duration              `json:"auto_complete_duration"`
	PickupAt             time.Time                  `json:"pickup_at"`
	PickupWindowDuration time.Duration              `json:"pickup_window_duration"`
	PrepTimeDuration     time.Duration              `json:"prep_time_duration"`
	Note                 string                     `json:"note"`
	PlacedAt             time.Time                  `json:"placed_at"`
	AcceptedAt           time.Time                  `json:"accepted_at"`
	RejectedAt           time.Time                  `json:"rejected_at"`
	ReadyAt              time.Time                  `json:"ready_at"`
	ExpiredAt            time.Time                  `json:"expired_at"`
	PickedUpAt           time.Time                  `json:"picked_up_at"`
	CanceledAt           time.Time                  `json:"canceled_at"`
	CancelReason         string                     `json:"cancel_reason"`
}

func (o *OrderFulfillmentPickupDetails) MarshalJSON() ([]byte, error) {
	jsonType := orderFulfillmentPickupDetails{
		Recipient:            o.Recipient,
		ExpiresAt:            o.ExpiresAt,
		AutoCompleteDuration: string(int(o.AutoCompleteDuration.Seconds())) + "S",
		PickupAt:             o.PickupAt,
		PickupWindowDuration: string(int(o.PickupWindowDuration.Seconds())) + "S",
		PrepTimeDuration:     string(int(o.PrepTimeDuration.Seconds())) + "S",
		Note:                 o.Note,
		PlacedAt:             o.PlacedAt,
		AcceptedAt:           o.AcceptedAt,
		RejectedAt:           o.RejectedAt,
		ReadyAt:              o.ReadyAt,
		ExpiredAt:            o.ExpiredAt,
		PickedUpAt:           o.PickedUpAt,
		CanceledAt:           o.CanceledAt,
		CancelReason:         o.CancelReason,
	}
	return json.Marshal(&jsonType)
}

func (o *OrderFulfillmentPickupDetails) UnmarshalJSON(b []byte) error {
	jsonType := orderFulfillmentPickupDetails{}
	err := json.Unmarshal(b, &jsonType)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails")
	}

	o.Recipient = jsonType.Recipient
	o.ExpiresAt = jsonType.ExpiresAt
	o.PickupAt = jsonType.PickupAt
	o.Note = jsonType.Note
	o.PlacedAt = jsonType.PlacedAt
	o.AcceptedAt = jsonType.AcceptedAt
	o.RejectedAt = jsonType.RejectedAt
	o.ReadyAt = jsonType.ReadyAt
	o.ExpiredAt = jsonType.ExpiredAt
	o.PickedUpAt = jsonType.PickedUpAt
	o.CanceledAt = jsonType.CanceledAt
	o.CancelReason = jsonType.CancelReason

	o.AutoCompleteDuration, err = duration.Parse(jsonType.AutoCompleteDuration)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.AutoCompleteDuration")
	}

	o.PickupWindowDuration, err = duration.Parse(jsonType.PickupWindowDuration)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.PickupWindowDuration")
	}

	o.PrepTimeDuration, err = duration.Parse(jsonType.PrepTimeDuration)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.PrepTimeDuration")
	}
	return nil
}
