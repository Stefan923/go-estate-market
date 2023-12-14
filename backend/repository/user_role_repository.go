package repository

import (
	"backend/data/database"
	"backend/data/model"
)

type UserRoleRepository struct {
	BaseRepository[model.UserRole]
}

func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{
		BaseRepository: BaseRepository[model.UserRole]{
			Database: database.GetDatabase(),
			Preloads: []preload{},
		},
	}
}
