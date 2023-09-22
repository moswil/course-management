package repository

import (
	"context"
	"log"

	"github.com/moswil/course-management/internal/core/dto"
	"github.com/moswil/course-management/internal/core/interface/repository"
	"github.com/moswil/course-management/internal/core/model"

	"gorm.io/gorm"
)

// MySQLCourseRepository implements the CourseRepository interface for MySQL.
type MySQLCourseRepository struct {
	db *gorm.DB // Initialize the GORM database connection here
}

// NewMySQLCourseRepository creates a new MySQLCourseRepository instance.
func NewMySQLCourseRepository(db *gorm.DB) repository.CourseRepository {
	// Initialize and inject the GORM database connection.
	return &MySQLCourseRepository{db: db}
}

// Create inserts a new course into the MySQL database using CreateCourseDTO.
func (r *MySQLCourseRepository) CreateCourse(ctx context.Context, createDTO *dto.CreateCourseDTO) (*dto.CourseDTO, error) {
	// Create a new Course model based on the CreateCourseDTO.
	course := model.NewCourse(createDTO.Title, createDTO.Instructor)

	log.Printf("course: %v: %v: %v\n", course.GetCourseID(), course.GetTitle(), course.GetInstructor())

	courseModel := &CourseModel{}
	courseModel.FromEntity(course)

	// Insert the course into MySQL.
	if err := r.db.WithContext(ctx).Create(courseModel).Error; err != nil {
		return nil, err
	}

	// return (*dto.CourseDTO)(courseModel), nil
	return courseModel.ToDTO(), nil
}

func (r *MySQLCourseRepository) GetCourseByID(ctx context.Context, courseID string) (*dto.CourseDTO, error) {
	// Create an empty Course model to hold the retrieved data.
	courseModel := &CourseModel{}

	if err := r.db.WithContext(ctx).Where("course_id = ?", courseID).First(&courseModel).Error; err != nil {
		// log the error
		return nil, err
	}
	return &dto.CourseDTO{
		CourseID:   courseModel.CourseID,
		Title:      courseModel.Title,
		Instructor: courseModel.Instructor,
	}, nil
}
