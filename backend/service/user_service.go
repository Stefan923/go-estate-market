package service

import (
	"backend/data/model"
	"backend/repository"
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
