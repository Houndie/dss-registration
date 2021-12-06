package object

import (
	"context"
	"time"
)

type Client interface {
	SignedPut(filesize int64, filename string) (string, error)
	SignedGet(filename string) (string, error)
	Exists(ctx context.Context, filename string) (bool, error)
	Delete(ctx context.Context, filename string) error
}

const (
	PutSigningDuration = 5 * time.Minute
	GetSigningDuration = 5 * time.Minute
	PutMaxSize         = 8 * 1024 * 1024
)
