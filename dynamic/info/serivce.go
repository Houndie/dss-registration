package info

import (
	"context"
)

type DB interface {
	Ping(context.Context) error
}

type Service struct {
	db      DB
	version string
}

func NewService(db DB, version string) *Service {
	return &Service{
		db:      db,
		version: version,
	}
}
