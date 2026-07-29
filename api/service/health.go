package service 

import "context"

type Repository interface {
  Ping(ctx context.Context) error
}

type HealthService struct {
  repo Repository
}

func NewHealthService(repo Repository) *HealthService {
  return &HealthService{repo: repo}
}

func (s *HealthService) Check(ctx context.Context) error {
  return s.repo.Ping(ctx)
}
