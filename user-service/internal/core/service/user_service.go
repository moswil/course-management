package service

import (
	"context"

	"github.com/moswil/course-management/user-service/internal/core/dto"
	"github.com/moswil/course-management/user-service/internal/core/port/repository"
	"github.com/moswil/course-management/user-service/internal/core/port/service"
)

type UserService struct {
	userRepository       repository.UserRepository
	oauthProviderService service.OAuthProviderService
}

func NewUserService(userRepo repository.UserRepository, ouathSvc service.OAuthProviderService) service.UserService {
	return &UserService{
		userRepository:       userRepo,
		oauthProviderService: ouathSvc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, createUser *dto.CreateUserDTO) (*dto.UserDTO, error) {
	// create user
	// user, err := s.userRepository.CreateUser(ctx, createUser)
	// if err != nil {
	// 	return nil, err
	// }

	// _ = user
	// call the oauth provider service
	userId, err := s.oauthProviderService.CreateUserInProvider(ctx, createUser)
	if err != nil {
		return nil, err
	}
	return &dto.UserDTO{
		UserId: userId,
	}, nil
}
