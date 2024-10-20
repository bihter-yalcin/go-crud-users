package service

import (
	models "api/model"
	"api/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.Repo.GetAllUsers()
}

func (service *UserService) GetUserByID(id int) (*models.User, error) {
	return service.Repo.GetUserByID(id)
}

func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	err := service.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) UpdateUser(updatedUser *models.User) (*models.User, error) {
	err := service.Repo.UpdateUser(updatedUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (service UserService) DeleteUser(id int) error {
	return service.Repo.DeleteUser(id)
}
