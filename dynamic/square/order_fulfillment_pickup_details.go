package square

import (
	"encoding/json"
	"time"

	"github.com/peterhellberg/duration"
	"github.com/pkg/errors"
)

type orderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient `json:"recipient,omitempty"`
	ExpiresAt            *time.Time                 `json:"expires_at,omitempty"`
	AutoCompleteDuration string                     `json:"auto_complete_duration,omitempty"`
	PickupAt             *time.Time                 `json:"pickup_at,omitempty"`
	PickupWindowDuration string                     `json:"pickup_window_duration,omitempty"`
	PrepTimeDuration     string                     `json:"prep_time_duration,omitempty"`
	Note                 string                     `json:"note,omitempty"`
	PlacedAt             *time.Time                 `json:"placed_at,omitempty"`
	AcceptedAt           *time.Time                 `json:"accepted_at,omitempty"`
	RejectedAt           *time.Time                 `json:"rejected_at,omitempty"`
	ReadyAt              *time.Time                 `json:"ready_at,omitempty"`
	ExpiredAt            *time.Time                 `json:"expired_at,omitempty"`
	PickedUpAt           *time.Time                 `json:"picked_up_at,omitempty"`
	CanceledAt           *time.Time                 `json:"canceled_at,omitempty"`
	CancelReason         string                     `json:"cancel_reason,omitempty"`
}

type OrderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient
	ExpiresAt            *time.Time
	AutoCompleteDuration *time.Duration
	PickupAt             *time.Time
	PickupWindowDuration *time.Duration
	PrepTimeDuration     *time.Duration
	Note                 string
	PlacedAt             *time.Time
	AcceptedAt           *time.Time
	RejectedAt           *time.Time
	ReadyAt              *time.Time
	ExpiredAt            *time.Time
	PickedUpAt           *time.Time
	CanceledAt           *time.Time
	CancelReason         string
}

func (o *OrderFulfillmentPickupDetails) MarshalJSON() ([]byte, error) {
	jsonType := orderFulfillmentPickupDetails{
		Recipient:    o.Recipient,
		ExpiresAt:    o.ExpiresAt,
		PickupAt:     o.PickupAt,
		Note:         o.Note,
		PlacedAt:     o.PlacedAt,
		AcceptedAt:   o.AcceptedAt,
		RejectedAt:   o.RejectedAt,
		ReadyAt:      o.ReadyAt,
		ExpiredAt:    o.ExpiredAt,
		PickedUpAt:   o.PickedUpAt,
		CanceledAt:   o.CanceledAt,
		CancelReason: o.CancelReason,
	}
	if o.AutoCompleteDuration != nil {
		jsonType.AutoCompleteDuration = string(int(o.AutoCompleteDuration.Seconds())) + "S"
	}
	if o.PickupWindowDuration != nil {
		jsonType.PickupWindowDuration = string(int(o.PickupWindowDuration.Seconds())) + "S"
	}
	if o.PrepTimeDuration != nil {
		jsonType.PrepTimeDuration = string(int(o.PrepTimeDuration.Seconds())) + "S"
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

	if jsonType.AutoCompleteDuration != "" {
		d, err := duration.Parse(jsonType.AutoCompleteDuration)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.AutoCompleteDuration")
		}
		o.AutoCompleteDuration = &d
	}

	if jsonType.PickupWindowDuration != "" {
		d, err := duration.Parse(jsonType.PickupWindowDuration)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.PickupWindowDuration")
		}
		o.PickupWindowDuration = &d
	}

	if jsonType.PrepTimeDuration != "" {
		d, err := duration.Parse(jsonType.PrepTimeDuration)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling OrderFulfillmentPickupDetails.PrepTimeDuration")
		}
		o.PrepTimeDuration = &d
	}
	return nil
}