package dto

type PropertyDto struct {
	Id              uint              `json:"id"`
	Owner           UserDetail        `json:"owner"`
	CurrentCurrency string            `json:"currentCurrency"`
	CurrentPrice    float32           `json:"currentPrice"`
	PropertyDetail  PropertyDetailDto `json:"propertyDetail"`
	Location        LocationDto       `json:"location"`
	Announcement    PostDto           `json:"announcement"`
}

type PropertyCreationDto struct {
	OwnerId           uint                      `json:"ownerId"`
	CurrentCurrencyId uint                      `json:"currentCurrencyId"`
	CurrentPrice      float32                   `json:"currentPrice"`
	CityId            uint                      `json:"cityId"`
	CategoryId        uint                      `json:"categoryId"`
	PropertyDetail    PropertyDetailCreationDto `json:"propertyDetail"`
	Announcement      PostCreationDto           `json:"announcement"`
}

type PropertyDetailDto struct {
	Id                    uint `json:"id"`
	NumberOfRooms         int  `json:"numberOfRooms"`
	NumberOfBathrooms     int  `json:"numberOfBathrooms"`
	NumberOfKitchens      int  `json:"numberOfKitchens"`
	NumberOfParkingSpaces int  `json:"numberOfParkingSpaces"`
}

type PropertyDetailCreationDto struct {
	NumberOfRooms         int `json:"numberOfRooms"`
	NumberOfBathrooms     int `json:"numberOfBathrooms"`
	NumberOfKitchens      int `json:"numberOfKitchens"`
	NumberOfParkingSpaces int `json:"numberOfParkingSpaces"`
}
