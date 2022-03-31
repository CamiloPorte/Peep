package seed

import (
	"context"
	peep "peep/pkg"
)

type service struct {
	repository peep.Repository
}

// Service provides db setup operations.
type Service interface {
	SetUpDB(ctx context.Context) error
	SeedDB(ctx context.Context) error
}

func NewService(repository peep.Repository) Service {
	return &service{repository}
}

func (s *service) SetUpDB(ctx context.Context) error {
	return s.repository.SetUpDB(ctx)
}

func (s *service) SeedDB(ctx context.Context) error {
	return s.repository.SeedDB(ctx)
}
