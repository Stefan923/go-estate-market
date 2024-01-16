package repository

import (
	db "github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
)

type UserAccountRepository struct {
	BaseRepository[model.UserAccount]
}

func NewUserAccountRepository() *UserAccountRepository {
	return &UserAccountRepository{
		BaseRepository: BaseRepository[model.UserAccount]{
			Database: db.GetDatabase(),
			Preloads: []preload{
				{string: "User"},
			},
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
		return nil, err
	}

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
		return false, err
	}

	return exists, err
}
