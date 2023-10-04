package service

import (
	"context"

	"github.com/moswil/course-management/user-service/internal/core/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (*dto.UserDTO, error)
}
