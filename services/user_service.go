package services

import (
	"compartamos-backend/models"
	"compartamos-backend/repositories"
	"errors"
)

type UserService struct {
    UserRepository repositories.UserRepository
}

func (service *UserService) CreateUser(user models.User) (models.User, error) {
    return service.UserRepository.CreateUser(user)
}

func (service *UserService) GetUsers() ([]models.User, error) {
    return service.UserRepository.GetUsers()
}

func (service *UserService) GetUser(id string) (models.User, error) {
    return service.UserRepository.GetUser(id)
}

func (service *UserService) UpdateUser(id string, user models.User) (models.User, error) {
    return service.UserRepository.UpdateUser(id, user)
}

func (service *UserService) DeleteUser(id string) error {
    user, err := service.UserRepository.GetUser(id)
    if err != nil {
        return errors.New("Usuario no encontrado")
    }

    if user.Age < 80 {
        return errors.New("Solo se pueden eliminar usuario mayores de 80 años")
    }

    return service.UserRepository.DeleteUser(id)
}
