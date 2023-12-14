package service

import (
	"backend/api/dto"
	"backend/config"
	"backend/data/model"
	error2 "backend/error"
	"backend/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserAccountService struct {
	userAccountRepository *repository.UserAccountRepository
	userService           *UserService
	tokenService          *TokenService
	roleService           *RoleService
	userRoleService       *UserRoleService
	config                *config.Config
}

func (service UserAccountService) Login(request dto.LoginRequest) (*dto.TokenDetail, error) {
	userAccount, err := service.userAccountRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	user, err := service.userService.GetById(userAccount.Id)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAccount.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	var roles []string
	for _, role := range *user.UserRoles {
		roles = append(roles, role.Role.Name)
	}

	var tokenDetail *dto.TokenDetail
	tokenDetail, err = service.tokenService.GenerateToken(&dto.TokenRequest{
		UserId: userAccount.Id,
		Email:  userAccount.Email,
		Roles:  roles,
	})

	return tokenDetail, nil
}

func (service UserAccountService) Register(context context.Context, request dto.RegisterRequest) (*dto.TokenDetail, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), service.config.AuthConfig.BCryptCost)
	if err != nil {
		return nil, err
	}

	if exists, _ := service.userAccountRepository.ExistsByEmail(request.Email); !exists {
		return nil, &error2.ServiceError{EndUserMessage: error2.EmailAlreadyUsed}
	}

	userAccount := model.UserAccount{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	createdUserAccount, err := service.userAccountRepository.Save(context, &userAccount)
	if err != nil {
		return nil, err
	}

	user := model.User{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		PhoneNumber:   request.PhoneNumber,
		Enabled:       false,
		UserAccountId: createdUserAccount.Id,
	}
	createdUser, err := service.userService.Save(context, &user)
	if err != nil {
		return nil, err
	}

	defaultRole, err := service.roleService.GetDefault()
	if err != nil {
		return nil, err
	}

	userRole := model.UserRole{
		UserId: createdUser.Id,
		RoleId: defaultRole.Id,
	}
	_, err = service.userRoleService.Save(context, &userRole)
	if err != nil {
		return nil, err
	}

	var roles []string
	roles = append(roles, defaultRole.Name)

	var tokenDetail *dto.TokenDetail
	tokenDetail, err = service.tokenService.GenerateToken(&dto.TokenRequest{
		UserId: createdUserAccount.Id,
		Email:  createdUserAccount.Email,
		Roles:  roles,
	})

	return tokenDetail, nil
}
