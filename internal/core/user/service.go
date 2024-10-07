package user

import "taskgo/pkg/exception"

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req CreateUserRequest) *exception.Exception {
	if err := s.repo.Create(req.ToUser()); err != nil {
		return err
	}
	return nil
}
