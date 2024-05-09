package service

import (
	"context"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Save(ctx context.Context, user model.User) (int, error) {
	return s.repo.Save(ctx, user)
}

func (s *UserService) GetUserByPhonenumber(phonenumber string) (model.User, error) {
	return s.repo.GetUserByPhonenumber(phonenumber)
}

func (s *UserService) GetAllWithRole(roleId int) ([]model.UserWithInfo, error) {
	return s.repo.GetAllWithRole(roleId)
}

func (s *UserService) Update(user model.User) error {
	return s.repo.Update(user)
}
