package service_test

import (
	"go-people-api/internal/models"
	"go-people-api/internal/service"
	"go-people-api/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockStore := new(mocks.UserRepository)
	user := models.User{
		ID:          1,
		Name:        "Dmitriy",
		Surname:     "Ushakov",
		Patronymic:  "Vasilevich",
		Age:         45,
		Gender:      "Male",
		Nationality: "RU",
	}
	mockStore.On("Create", user).Return(nil).Once()

	service := service.NewUserService(mockStore)

	err := service.Create(user)

	assert.Nil(t, err)

}

func TestGet(t *testing.T) {
	mockStore := new(mocks.UserRepository)
	user := models.User{
		ID:          1,
		Name:        "Dmitriy",
		Surname:     "Ushakov",
		Patronymic:  "Vasilevich",
		Age:         45,
		Gender:      "Male",
		Nationality: "RU",
	}
	mockStore.On("Get", "Dmitriy").Return(user, nil).Once()
	service := service.NewUserService(mockStore)
	result, err := service.Get("Dmitriy")

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockStore.AssertExpectations(t)

}
