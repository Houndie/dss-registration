package square

import "time"

type BusinessHours struct {
	Periods []*BusinessHoursPeriod `json:"periods"`
}

type BusinessHoursPeriod struct {
	DayOfWeek      string    `json:"day_of_week"`
	StartLocalTime time.Time `json:"start_local_time"`
	EndLocalTime   time.Time `json:"end_local_time"`
}
