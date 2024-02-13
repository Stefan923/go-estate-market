package dto

type PropertyDto struct {
	Owner           UserDetail        `json:"owner"`
	CurrentCurrency string            `json:"currentCurrency"`
	CurrentPrice    float32           `json:"currentPrice"`
	PropertyDetail  PropertyDetailDto `json:"propertyDetail"`
	Location        LocationDto       `json:"location"`
	Announcement    PostDto           `json:"announcement"`
}

type PropertyCreationDto struct {
	OwnerId         uint    `json:"ownerId"`
	CurrentCurrency string  `json:"currentCurrency"`
	CurrentPrice    float32 `json:"currentPrice"`
}

type PropertyDetailDto struct {
	NumberOfRooms         int `json:"numberOfRooms"`
	NumberOfBathrooms     int `json:"numberOfBathrooms"`
	NumberOfKitchens      int `json:"numberOfKitchens"`
	NumberOfParkingSpaces int `json:"numberOfParkingSpaces"`
}
