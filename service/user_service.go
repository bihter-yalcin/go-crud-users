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
