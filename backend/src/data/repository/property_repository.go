package repository

import (
	"fmt"
	"github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/pagination"
)

type PropertyRepository struct {
	BaseRepository[model.Property]
}

func NewPropertyRepository() *PropertyRepository {
	return &PropertyRepository{
		BaseRepository: BaseRepository[model.Property]{
			Database: database.GetDatabase(),
			Preloads: []PreloadSetting{
				{EntityName: "PropertyCategory"},
				{EntityName: "Owner"},
				{EntityName: "Owner.UserAccount"},
				{EntityName: "City"},
				{EntityName: "City.State"},
				{EntityName: "City.State.Country"},
				{EntityName: "Announcement"},
				{EntityName: "CurrentCurrency"},
			},
		},
	}
}

func (repository *PropertyRepository) FindAllByCategory(category string, pageInfo *pagination.PageInfo) (*pagination.Page[model.Property], error) {
	var properties []model.Property
	var preloadedDatabase = Preload(repository.Database, repository.Preloads)

	offset := (pageInfo.PageNumber - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize

	err := preloadedDatabase.Where("property_category_id = ?", category).
		Order(fmt.Sprintf("%s %s", pageInfo.SortBy, pageInfo.SortType)).
		Offset(offset).
		Limit(limit).
		Find(&properties).Error
	if err != nil {
		return nil, err
	}

	return &pagination.Page[model.Property]{
		Elements:   properties,
		PageNumber: pageInfo.PageNumber,
		PageSize:   pageInfo.PageSize,
	}, nil
}
