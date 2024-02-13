package model

type Post struct {
	BaseModel
	Title        string `gorm:"type:string; size:64; not null;"`
	Description  string `gorm:"type:string; size:1024; not null;"`
	PostComments *[]PostComment
	PropertyId   uint `gorm:"unique;"`
}

type PostComment struct {
	BaseModel
	Description string `gorm:"type:string; size:512; not null;"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	PostId      uint
	Post        Post `gorm:"foreignKey:PostId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}
