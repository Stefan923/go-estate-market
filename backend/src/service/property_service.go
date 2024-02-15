package service

import (
	"context"
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
	postService              *PostService
}

func NewPropertyService() *PropertyService {
	propertyRepository := repository.NewPropertyRepository()
	return &PropertyService{
		BaseService: BaseService[model.Property, model.Property, model.Property, model.Property]{
			Repository: &propertyRepository.BaseRepository,
		},
		propertyRepository:       propertyRepository,
		propertyDetailRepository: repository.NewPropertyDetailRepository(),
		postService:              NewPostService(),
	}
}

func (service *PropertyService) Save(context context.Context, propertyDto dto.PropertyCreationDto) (*dto.PropertyDto, error) {
	property := model.Property{
		OwnerId:            propertyDto.OwnerId,
		CityId:             propertyDto.CityId,
		CurrentCurrency:    propertyDto.CurrentCurrency,
		CurrentPrice:       propertyDto.CurrentPrice,
		PropertyCategoryId: propertyDto.CategoryId,
	}

	createdProperty, err := service.propertyRepository.Save(context, &property)
	if err != nil {
		return nil, err
	}

	propertyDetail := model.PropertyDetail{
		NumberOfRooms:         propertyDto.PropertyDetail.NumberOfRooms,
		NumberOfBathrooms:     propertyDto.PropertyDetail.NumberOfBathrooms,
		NumberOfKitchens:      propertyDto.PropertyDetail.NumberOfKitchens,
		NumberOfParkingSpaces: propertyDto.PropertyDetail.NumberOfParkingSpaces,
		PropertyId:            createdProperty.Id,
	}

	createdPropertyDetail, err := service.propertyDetailRepository.Save(context, &propertyDetail)
	if err != nil {
		return nil, err
	}

	post := dto.PostCreationWithIdDto{
		Title:       propertyDto.Announcement.Title,
		Description: propertyDto.Announcement.Description,
		PropertyId:  createdProperty.Id,
	}

	createdPost, err := service.postService.Save(context, &post)

	createdPropertyDto := dto.PropertyDto{
		Id: property.Id,
		Owner: dto.UserDetail{
			Email:     createdProperty.Owner.UserAccount.Email,
			FirstName: createdProperty.Owner.FirstName,
			LastName:  createdProperty.Owner.LastName,
		},
		CurrentCurrency: createdProperty.CurrentCurrency,
		CurrentPrice:    createdProperty.CurrentPrice,
		PropertyDetail: dto.PropertyDetailDto{
			NumberOfRooms:         createdPropertyDetail.NumberOfRooms,
			NumberOfBathrooms:     createdPropertyDetail.NumberOfBathrooms,
			NumberOfKitchens:      createdPropertyDetail.NumberOfKitchens,
			NumberOfParkingSpaces: createdPropertyDetail.NumberOfParkingSpaces,
		},
		Location: dto.LocationDto{
			Country: createdProperty.City.State.Country.Name,
			State:   createdProperty.City.State.Name,
			City:    createdProperty.City.Name,
		},
		Announcement: dto.PostDto{
			Id:          createdPost.Id,
			Title:       createdPost.Title,
			Description: createdPost.Description,
			PropertyId:  createdPost.PropertyId,
		},
	}

	return &createdPropertyDto, nil
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
