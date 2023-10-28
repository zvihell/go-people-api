package service

import "go-people-api/internal/models"

type UserService interface {
	Create(user models.User) error
	Get(name string) (models.User, error)
	Delete(id int) error
	Update(id int, user models.User) error
}

type Service struct {
	userRepo UserService
}

func NewUserService(userRepo UserService) *Service {
	return &Service{userRepo: userRepo}
}

func (s *Service) Create(user models.User) error {
	return s.userRepo.Create(user)

}

func (s *Service) Get(name string) (models.User, error) {
	return s.userRepo.Get(name)

}

func (s *Service) Delete(id int) error {
	return s.userRepo.Delete(id)

}

func (s *Service) Update(id int, user models.User) error {
	return s.userRepo.Update(id, user)
}
