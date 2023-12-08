package model

import "time"

type Property struct {
	BaseModel
	OwnerId              uint
	Owner                User `gorm:"foreignKey:OwnerId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	CityId               uint
	City                 User `gorm:"foreignKey:CityId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	PropertyPrices       *[]PropertyPrice
	PropertyCategoryId   uint
	PropertyCategory     PropertyCategory `gorm:"foreignKey:PropertyCategoryId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	PropertyAnnouncement Post             `gorm:"foreignKey:PropertyId;"`
}

type PropertyCategory struct {
	BaseModel
	Name       string `gorm:"size:64; type:string; not null; unique;"`
	Icon       string `gorm:"size:128; type:string; not null; unique;"`
	Properties *[]Property
}

type PropertyDetails struct {
	BaseModel
	NumberOfRooms     int `gorm:"type:int; not null;"`
	NumberOfBathrooms int `gorm:"type:int; not null;"`
	NumberOfKitchens  int `gorm:"type:int; not null;"`
	PropertyId        uint
	Property          Property `gorm:"foreignKey:PropertyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}

type PropertyPrice struct {
	BaseModel
	Price      float32   `gorm:"type:decimal(10,2); not null;"`
	PriceAt    time.Time `gorm:"type:TIMESTAMP with time zone; not null;"`
	PropertyId uint
	Property   Property `gorm:"foreignKey:PropertyId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}
