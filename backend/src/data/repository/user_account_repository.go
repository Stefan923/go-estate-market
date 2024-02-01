package repository

import (
	db "github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/metrics"
	"reflect"
)

type UserAccountRepository struct {
	BaseRepository[model.UserAccount]
}

func NewUserAccountRepository() *UserAccountRepository {
	return &UserAccountRepository{
		BaseRepository: BaseRepository[model.UserAccount]{
			Database: db.GetDatabase(),
		},
	}
}

func (repository UserAccountRepository) FindByEmail(email string) (*model.UserAccount, error) {
	userAccount := new(model.UserAccount)
	database := Preload(repository.Database, repository.Preloads)

	err := database.
		Where("email = ? and deleted_at is null", email).
		First(userAccount).
		Error
	if err != nil {
		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*userAccount).String(), "FindByEmail", "Failed").Inc()
		return nil, err
	}

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*userAccount).String(), "FindByEmail", "Success").Inc()
	return userAccount, nil
}

func (repository UserAccountRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool

	err := repository.Database.Model(&model.UserAccount{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error
	if err != nil {
		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(model.UserAccount{}).String(), "FindByEmail", "Failed").Inc()
		return false, err
	}

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(model.UserAccount{}).String(), "FindByEmail", "Success").Inc()
	return exists, err
}
