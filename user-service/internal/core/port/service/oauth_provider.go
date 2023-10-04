package service

import (
	"context"

	"github.com/moswil/course-management/user-service/internal/core/dto"
)

type OAuthProviderService interface {
	// CreateUserInProvider
	CreateUserInProvider(ctx context.Context, user *dto.CreateUserDTO) (userId string, err error)
}
