package repository

import (
	"backend/data/database"
	"backend/data/model"
)

type RoleRepository struct {
	BaseRepository[model.Role]
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		BaseRepository: BaseRepository[model.Role]{
			Database: database.GetDatabase(),
			Preloads: []preload{},
		},
	}
}

func (repository *RoleRepository) FindDefault() (*model.Role, error) {
	role := new(model.Role)

	err := repository.Database.
		Where("default = ? and deleted_at is null", true).
		First(role).
		Error
	if err != nil {
		return nil, err
	}

	return role, nil
}
