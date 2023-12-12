package repository

import (
	db "backend/data/database"
	"backend/data/model"
	"backend/util"
	"context"
	"database/sql"
	"gorm.io/gorm"
	"log"
	"time"
)

const userIdFiledName = "UserId"

type preload struct {
	string
}

type BaseRepository[T any] struct {
	Database *gorm.DB
	Preloads []preload
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{
		Database: db.GetDatabase(),
	}
}

func (repository *BaseRepository[T]) FindById(id uint) (*T, error) {
	object := new(T)
	database := Preload(repository.Database, repository.Preloads)

	err := database.
		Where("id = ? and deleted_at is null", id).
		First(object).
		Error
	if err != nil {
		log.Println("")
		return nil, err
	}

	return object, nil
}

func (repository *BaseRepository[T]) Save(context context.Context, object *T) (*T, error) {
	database := repository.Database.WithContext(context).Begin()

	err := database.
		Create(object).
		Error
	if err != nil {
		database.Rollback()
		log.Println("Error while creating object: ", err)
		return nil, err
	}
	database.Commit()

	convertedObject, _ := util.ConvertTo[model.BaseModel](object)
	return repository.FindById(convertedObject.Id)
}

func (repository *BaseRepository[T]) Update(context context.Context, id uint, object *T) (*T, error) {
	updateObject, _ := util.ConvertTo[model.BaseModel](object)
	objectMap, _ := util.ConvertToSnakeCaseMap(updateObject)
	objectMap["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	updatedObject := new(T)
	database := repository.Database.WithContext(context).Begin()

	err := database.Model(updatedObject).
		Where("id = ? and deleted_at is null", id).
		Updates(objectMap).
		Error
	if err != nil {
		database.Rollback()
		return nil, err
	}
	database.Commit()

	return repository.FindById(id)
}

func (repository *BaseRepository[T]) Delete(context context.Context, id uint) error {
	database := repository.Database.WithContext(context).Begin()

	object := new(T)
	deleteMap := map[string]interface{}{
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	if context.Value(userIdFiledName) == nil {
		return nil
	}
	if deletedObjectsCount := database.
		Model(object).
		Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).
		RowsAffected; deletedObjectsCount == 0 {
		database.Rollback()
		return nil
	}
	database.Commit()

	return nil
}

func Preload(database *gorm.DB, preloads []preload) *gorm.DB {
	for _, preload := range preloads {
		database = database.Preload(preload.string)
	}
	return database
}
