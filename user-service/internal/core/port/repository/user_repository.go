package repository

import (
	"context"

	"github.com/moswil/course-management/user-service/internal/core/dto"
)

type UserRepository interface {
	// CreateUser
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (*dto.UserDTO, error)
}
