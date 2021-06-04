package info

import "context"

type Healthiness string

const (
	Healthy   Healthiness = "healthy"
	Unhealthy Healthiness = "unhealthy"
)

func (s *Service) Health(ctx context.Context) Healthiness {
	if err := s.db.Ping(ctx); err != nil {
		return Unhealthy
	}

	return Healthy
}
