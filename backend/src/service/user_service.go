package service

import (
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/repository"
)

type UserService struct {
	BaseService[model.User, model.User, model.User, model.User]
}

func NewUserService() *UserService {
	return &UserService{
		BaseService: BaseService[model.User, model.User, model.User, model.User]{
			Repository: repository.NewBaseRepository[model.User](),
		},
	}
}
