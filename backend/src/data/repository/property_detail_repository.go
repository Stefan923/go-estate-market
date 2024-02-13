package repository

import (
	"github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/metrics"
	"reflect"
)

type PropertyDetailRepository struct {
	BaseRepository[model.PropertyDetail]
}

func NewPropertyDetailRepository() *PropertyDetailRepository {
	return &PropertyDetailRepository{
		BaseRepository: BaseRepository[model.PropertyDetail]{
			Database: database.GetDatabase(),
			Preloads: []PreloadSetting{},
		},
	}
}

func (repository *PropertyDetailRepository) FindByPropertyId(propertyId uint) (*model.PropertyDetail, error) {
	var propertyDetail *model.PropertyDetail

	err := repository.Database.
		Where("property_id = ? and deleted_at is null", propertyId).
		First(propertyDetail).
		Error
	if err != nil {
		metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*propertyDetail).String(), "FindByPropertyId", "Failed").Inc()
		return nil, err
	}

	metrics.DatabaseCallCounter.WithLabelValues(reflect.TypeOf(*propertyDetail).String(), "FindByPropertyId", "Success").Inc()
	return propertyDetail, nil
}
