package storage

import "time"

type Volunteer struct {
	ID        string
	UserID    string
	Name      string
	Email     string
	Disabled  bool
	CreatedAt time.Time
}
