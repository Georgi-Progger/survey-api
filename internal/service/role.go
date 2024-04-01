package service

import (
	"context"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type RoleService struct {
	repo repository.Role
}

func NewRoleService(repo repository.Role) *RoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) GetByName(ctx context.Context, name string) (*model.Role, error) {
	return r.repo.GetByName(ctx, name)
}

func (r *RoleService) SetRole(userId, roleId int) error {
	return r.repo.SetRole(userId,roleId)
}