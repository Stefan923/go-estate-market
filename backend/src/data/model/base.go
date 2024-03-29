package model

type Country struct {
	BaseModel
	Name   string `gorm:"size:15; type:string; not null;"`
	States *[]State
}

type State struct {
	BaseModel
	Name      string `gorm:"size:15; type:string; not null;"`
	CountryId uint
	Country   Country `gorm:"foreignKey:CountryId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	Cities    *[]City
}

type City struct {
	BaseModel
	Name       string `gorm:"size:15; type:string; not null;"`
	StateId    uint
	State      State `gorm:"foreignKey:StateId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	Properties *[]Property
}

type Address struct {
	BaseModel
	StreetName   string `gorm:"size:63; type:string; not null;"`
	StreetNumber int    `gorm:"type:int; not null;"`
	Floor        int    `gorm:"type:int; not null;"`
	Apartment    int    `gorm:"type:int; not null;"`
}

type Currency struct {
	BaseModel
	Code string `gorm:"size:4; type:string; not null;"`
}
