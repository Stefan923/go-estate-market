package service

import (
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/pagination"
	"github.com/Stefan923/go-estate-market/data/repository"
)

type PropertyService struct {
	BaseService[model.Property, model.Property, model.Property, model.Property]
	propertyRepository *repository.PropertyRepository
}

func NewPropertyService() *PropertyService {
	propertyRepository := repository.NewPropertyRepository()
	return &PropertyService{
		BaseService: BaseService[model.Property, model.Property, model.Property, model.Property]{
			Repository: &propertyRepository.BaseRepository,
		},
		propertyRepository: propertyRepository,
	}
}

func (service *PropertyService) GetAllByCategory(category string, pageInfo *pagination.PageInfo) (*pagination.Page[model.Property], error) {
	return service.propertyRepository.FindAllByCategory(category, pageInfo)
}
