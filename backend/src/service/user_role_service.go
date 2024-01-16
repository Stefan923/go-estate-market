package service

import (
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/repository"
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
