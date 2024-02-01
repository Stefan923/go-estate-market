package repository

import (
	"github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/metrics"
	"reflect"
)

type RoleRepository struct {
	BaseRepository[model.Role]
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		BaseRepository: BaseRepository[model.Role]{
			Database: database.GetDatabase(),
			Preloads: []PreloadSetting{},
		},
	}
}

func (repository *RoleRepository) FindDefault() (*model.Role, error) {
	role := new(model.Role)

	err := repository.Database.
		Where("default_role = ? and deleted_at is null", true).
		First(role).
		Error
	if err != nil {
		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*role).String(), "FindDefault", "Failed").Inc()
		return nil, err
	}

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*role).String(), "FindDefault", "Success").Inc()
	return role, nil
}
