package repository

import (
	"context"
	"database/sql"
	db "github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/metrics"
	"github.com/Stefan923/go-estate-market/util"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
)

const userIdFiledName = "UserId"

type PreloadSetting struct {
	EntityName string
}

type BaseRepository[T any] struct {
	Database *gorm.DB
	Preloads []PreloadSetting
}

func NewBaseRepository[T any](preloads []PreloadSetting) *BaseRepository[T] {
	return &BaseRepository[T]{
		Database: db.GetDatabase(),
		Preloads: preloads,
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
		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "FindById", "Failed").Inc()
		return nil, err
	}

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "FindById", "Success").Inc()
	return object, nil
}

func (repository *BaseRepository[T]) Save(context context.Context, object *T) (*T, error) {
	database := repository.Database.WithContext(context).Begin()

	err := database.
		Create(object).
		Error
	if err != nil {
		database.Rollback()

		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Create", "Failed").Inc()
		log.Println("Error while creating object: ", err)
		return nil, err
	}
	database.Commit()

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Create", "Success").Inc()
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

		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Update", "Failed").Inc()
		return nil, err
	}
	database.Commit()

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Update", "Success").Inc()
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

		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Delete", "Failed").Inc()
		return nil
	}
	database.Commit()

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*object).String(), "Delete", "Success").Inc()
	return nil
}

func Preload(database *gorm.DB, preloads []PreloadSetting) *gorm.DB {
	for _, preload := range preloads {
		database = database.Preload(preload.EntityName)
	}
	return database
}
