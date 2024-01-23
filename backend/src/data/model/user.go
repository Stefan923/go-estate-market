package model

type User struct {
	BaseModel
	FirstName     string `gorm:"type:string; size:64; not null;"`
	LastName      string `gorm:"type:string; size:64; not null;"`
	PhoneNumber   string `gorm:"type:string; size:16; not null;"`
	Enabled       bool   `gorm:"type:bool; default:true;"`
	UserAccountId uint
	UserAccount   UserAccount `gorm:"foreignKey:UserAccountId; constraint:OnUpdate:NO ACTION; OnDelete: NO ACTION;"`
	UserRoles     *[]UserRole
}

type UserAccount struct {
	BaseModel
	Email    string `gorm:"type:string; size:64; not null; unique;"`
	Password string `gorm:"type:string; 256; not null;"`
}

type Role struct {
	BaseModel
	Name        string `gorm:"type:string; size:10; not null; unique;"`
	DefaultRole bool   `gorm:"type:bool; default:false;"`
	UserRoles   *[]UserRole
}

type UserRole struct {
	BaseModel
	UserId uint
	RoleId uint
	User   User `gorm:"foreignKey:UserId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
	Role   Role `gorm:"foreignKey:RoleId; constraint:OnUpdate:NO ACTION; OnDelete:NO ACTION;"`
}
