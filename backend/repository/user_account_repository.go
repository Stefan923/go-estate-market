package repository

import (
	db "backend/data/database"
	"backend/data/model"
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
		Where("email = ? and deleted_at is null").
		First(userAccount).
		Error
	if err != nil {
		return nil, err
	}

	return userAccount, nil
}
