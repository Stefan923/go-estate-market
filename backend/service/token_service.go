package service

import (
	"backend/api/dto"
	"backend/config"
	error2 "backend/error"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	UserIdKey     string = "UserId"
	EmailKey      string = "Email"
	RolesKey      string = "Roles"
	ExpireTimeKey string = "ExpireTime"
)

type TokenService struct {
	config *config.Config
}

func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}

func (service *TokenService) GenerateToken(token *dto.TokenRequest) (*dto.TokenDetail, error) {
	tokenDetail := new(dto.TokenDetail)
	tokenDetail.AccessTokenExpireTime = time.Now().Add(service.config.JWT.AccessTokenExpireDurationMinutes * time.Minute).Unix()
	tokenDetail.AccessTokenExpireTime = time.Now().Add(service.config.JWT.AccessTokenExpireDurationMinutes * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims[UserIdKey] = token.UserId
	accessTokenClaims[EmailKey] = token.Email
	accessTokenClaims[RolesKey] = token.Roles
	accessTokenClaims[ExpireTimeKey] = tokenDetail.AccessTokenExpireTime

	var err error
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(service.config.JWT.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[UserIdKey] = token.UserId
	refreshTokenClaims[ExpireTimeKey] = tokenDetail.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(service.config.JWT.RefreshTokenSecret))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func (service *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &error2.InternalError{EndUserMessage: error2.InvalidToken}
		}
		return []byte(service.config.JWT.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (service *TokenService) GetClaims(token string) (map[string]interface{}, error) {
	claimsMap := map[string]interface{}{}

	accessToken, err := service.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if ok && accessToken.Valid {
		for key, value := range claims {
			claimsMap[key] = value
		}
		return claimsMap, nil
	}

	return nil, &error2.InternalError{EndUserMessage: error2.ClaimsNotFound}
}
