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

func (service UserService) CreateUser(user *models.User) error {
	return service.Repo.CreateUser(user)
}

func (service UserService) UpdateUser(user *models.User) error {
	return service.Repo.UpdateUser(user)
}

func (service UserService) DeleteUser(id int) error {
	return service.Repo.DeleteUser(id)
}
