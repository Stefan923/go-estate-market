package service

import (
	"github.com/Stefan923/go-estate-market/api/dto"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/pagination"
	"github.com/Stefan923/go-estate-market/data/repository"
	error3 "github.com/Stefan923/go-estate-market/error"
)

type PropertyService struct {
	BaseService[model.Property, model.Property, model.Property, model.Property]
	propertyRepository       *repository.PropertyRepository
	propertyDetailRepository *repository.PropertyDetailRepository
}

func NewPropertyService() *PropertyService {
	propertyRepository := repository.NewPropertyRepository()
	return &PropertyService{
		BaseService: BaseService[model.Property, model.Property, model.Property, model.Property]{
			Repository: &propertyRepository.BaseRepository,
		},
		propertyRepository:       propertyRepository,
		propertyDetailRepository: repository.NewPropertyDetailRepository(),
	}
}

func (service *PropertyService) GetAllByCategory(category string, pageInfo *pagination.PageInfo) (*pagination.Page[dto.PropertyDto], error) {
	properties, err := service.propertyRepository.FindAllByCategory(category, pageInfo)
	if err != nil {
		return nil, &error3.InternalError{EndUserMessage: error3.RecordNotFound}
	}

	var mappedProperties *pagination.Page[dto.PropertyDto]
	for _, property := range *properties.Elements {
		propertyDetail, err := service.propertyDetailRepository.FindByPropertyId(property.Id)
		if err != nil {
			return nil, &error3.InternalError{EndUserMessage: error3.RecordNotFound}
		}

		mappedProperty := dto.PropertyDto{
			Owner: dto.UserDetail{
				Email:     property.Owner.UserAccount.Email,
				FirstName: property.Owner.FirstName,
				LastName:  property.Owner.LastName,
			},
			CurrentCurrency: property.CurrentCurrency,
			CurrentPrice:    property.CurrentPrice,
			PropertyDetail: dto.PropertyDetailDto{
				NumberOfRooms:         propertyDetail.NumberOfRooms,
				NumberOfBathrooms:     propertyDetail.NumberOfBathrooms,
				NumberOfKitchens:      propertyDetail.NumberOfKitchens,
				NumberOfParkingSpaces: propertyDetail.NumberOfParkingSpaces,
			},
			Location: dto.LocationDto{
				Country: property.City.State.Country.Name,
				State:   property.City.State.Name,
				City:    property.City.Name,
			},
			Announcement: dto.PostDto{
				Title:       property.Announcement.Title,
				Description: property.Announcement.Description,
				PropertyId:  property.Id,
			},
		}
		*mappedProperties.Elements = append(*mappedProperties.Elements, mappedProperty)
	}

	mappedProperties.PageNumber = properties.PageNumber
	mappedProperties.PageSize = properties.PageSize

	return mappedProperties, nil
}
