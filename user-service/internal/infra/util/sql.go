package util

import (
	"github.com/moswil/course-management/user-service/internal/core/port/repository"
	"github.com/moswil/course-management/user-service/internal/infra/repository/sql/dao"
	"github.com/moswil/course-management/user-service/internal/infra/repository/sql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQLRepository() (repository.UserRepository, error) {
	// Use config
	// Replace with your actual MySQL database connection details
	dsn := "codeguru:CodeGuru96@tcp(localhost:3306)/courses"

	// Initialize GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the Course model
	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		return nil, err
	}

	// Initialize the GORM repository
	userRepository := dao.NewMySQLUserRepository(db)

	return userRepository, nil
}
