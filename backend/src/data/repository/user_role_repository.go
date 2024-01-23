package repository

import (
	"github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
)

type UserRoleRepository struct {
	BaseRepository[model.UserRole]
}

func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{
		BaseRepository: BaseRepository[model.UserRole]{
			Database: database.GetDatabase(),
			Preloads: []PreloadSetting{
				{EntityName: "Role"},
				{EntityName: "User"},
			},
		},
	}
}
