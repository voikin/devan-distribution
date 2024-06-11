package user

import "context"

func (s *Service) CreateUser(ctx context.Context) error {
	return s.repo.CreateUser(ctx)
}
