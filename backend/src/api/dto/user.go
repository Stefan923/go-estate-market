package dto

type RegisterRequest struct {
	FirstName   string `json:"firstName" binding:"required,min=3"`
	LastName    string `json:"lastName" binding:"required,min=3"`
	PhoneNumber string `json:"phoneNumber" binding:"required,min=3"`
	Email       string `json:"email" binding:"required,min=6,email"`
	Password    string `json:"password" binding:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,min=6,email"`
	Password string `json:"password" binding:"required"`
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type TokenRequest struct {
	UserId uint
	Email  string
	Roles  []string
}
