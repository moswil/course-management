package dao

import (
	"context"

	"github.com/moswil/course-management/user-service/internal/core/dto"
	"github.com/moswil/course-management/user-service/internal/core/port/repository"
	"github.com/moswil/course-management/user-service/internal/infra/repository/sql/model"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB // Initialize the GORM database connection here
}

// NewMySQLUserRepository creates a new MySQLUserRepository instance.
func NewMySQLUserRepository(db *gorm.DB) repository.UserRepository {
	// Initialize and inject the GORM database connection.
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) CreateUser(ctx context.Context, user *dto.CreateUserDTO) (*dto.UserDTO, error) {
	userModel := &model.UserModel{
		UserId: "test",
		AuthId: "test",
	}
	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		UserId: "test",
	}, nil
}
