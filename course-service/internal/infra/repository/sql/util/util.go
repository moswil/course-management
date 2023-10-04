package util

import (
	"log"

	"github.com/moswil/course-management/internal/core/port/repository"
	"github.com/moswil/course-management/internal/infra/repository/sql/dao"
	"github.com/moswil/course-management/internal/infra/repository/sql/model"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQLRepository(tracer trace.Tracer) (repository.CourseRepository, error) {
	// Replace with your actual MySQL database connection details
	// dsn := "codeguru:CodeGuru96@tcp(localhost:3306)/courses"
	dsn := viper.GetString("database.mysql.dsn")

	log.Printf("MySQL DSN: %v\n", dsn)

	// Initialize GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// add telemetry (instrumentation)
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	// Auto-migrate the Course model
	if err := db.AutoMigrate(&model.CourseModel{}); err != nil {
		return nil, err
	}

	// Initialize the GORM repository
	courseRepository := dao.NewMySQLCourseRepository(db, tracer)

	return courseRepository, nil
}
