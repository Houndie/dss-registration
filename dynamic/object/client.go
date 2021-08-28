package object

import "time"

type Client interface {
	SignedPut(filesize int64, filename string) (string, error)
}

const (
	PutSigningDuration = 5 * time.Minute
	PutMaxSize         = 8 * 1024 * 1024
)
