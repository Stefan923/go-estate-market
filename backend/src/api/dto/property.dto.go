package dto

type PropertyDto struct {
	Owner           UserDetail `json:"owner"`
	CurrentCurrency string     `json:"currentCurrency"`
	CurrentPrice    float32    `json:"currentPrice"`
}

type PropertyDetailDto struct {
	NumberOfRooms         int `json:"numberOfRooms"`
	NumberOfBathrooms     int `json:"numberOfBathrooms"`
	NumberOfKitchens      int `json:"numberOfKitchens"`
	NumberOfParkingSpaces int `json:"numberOfParkingSpaces"`
}
