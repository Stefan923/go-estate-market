package service

import (
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/repository"
)

type RoleService struct {
	BaseService[model.Role, model.Role, model.Role, model.Role]
	roleRepository *repository.RoleRepository
}

func NewRoleService() *RoleService {
	roleRepository := repository.NewRoleRepository()
	return &RoleService{
		BaseService: BaseService[model.Role, model.Role, model.Role, model.Role]{
			Repository: &roleRepository.BaseRepository,
		},
		roleRepository: roleRepository,
	}
}

func (service *RoleService) GetDefault() (*model.Role, error) {
	return service.roleRepository.FindDefault()
}
