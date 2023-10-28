package repository

import "go-people-api/internal/models"

//go:generate go run github.com/vektra/mockery/v2@v2.36.0  --all
type UserRepository interface {
	Create(user models.User) error
	Get(name string) (models.User, error)
	Delete(id int) error
	Update(id int, user models.User) error
}
