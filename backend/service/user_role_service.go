package service

import (
	"backend/data/model"
	"backend/repository"
)

type UserRoleService struct {
	BaseService[model.UserRole, model.UserRole, model.UserRole, model.UserRole]
	userRoleRepository *repository.UserRoleRepository
}

func NewUserRoleService() *UserRoleService {
	userRoleRepository := repository.NewUserRoleRepository()
	return &UserRoleService{
		BaseService: BaseService[model.UserRole, model.UserRole, model.UserRole, model.UserRole]{
			Repository: &userRoleRepository.BaseRepository,
		},
		userRoleRepository: userRoleRepository,
	}
}
