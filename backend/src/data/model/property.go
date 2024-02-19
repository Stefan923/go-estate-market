package model

import (
	"time"
)

type Property struct {
	BaseModel
	OwnerId            uint
	Owner              User `gorm:"foreignKey:OwnerId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	CityId             uint
	City               City `gorm:"foreignKey:CityId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	CurrentCurrencyId  uint
	CurrentCurrency    Currency `gorm:"foreignKey:CurrentCurrencyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	CurrentPrice       float32  `gorm:"type:decimal(10,2); not null;"`
	Prices             *[]PropertyPrice
	PropertyCategoryId uint
	Category           PropertyCategory `gorm:"foreignKey:PropertyCategoryId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	Announcement       Post             `gorm:"foreignKey:PropertyId;"`
}

type PropertyCategory struct {
	BaseModel
	Name       string `gorm:"size:64; type:string; not null; unique;"`
	Icon       string `gorm:"size:128; type:string; not null; unique;"`
	Properties *[]Property
}

type PropertyDetail struct {
	BaseModel
	NumberOfRooms         int `gorm:"type:int; not null;"`
	NumberOfBathrooms     int `gorm:"type:int; not null;"`
	NumberOfKitchens      int `gorm:"type:int; not null;"`
	NumberOfParkingSpaces int `gorm:"type:int; default:0"`
	PropertyId            uint
	Property              Property `gorm:"foreignKey:PropertyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}

type PropertyPrice struct {
	BaseModel
	CurrencyId uint
	Currency   Currency  `gorm:"foreignKey:CurrencyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	Price      float32   `gorm:"type:decimal(10,2); not null;"`
	PriceAt    time.Time `gorm:"type:TIMESTAMP with time zone; not null;"`
	PropertyId uint
	Property   Property `gorm:"foreignKey:PropertyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}
