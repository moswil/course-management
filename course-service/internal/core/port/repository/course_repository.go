package repository

import (
	"context"

	"github.com/moswil/course-management/internal/core/dto"
)

// CourseRepository defines the methods for interacting with the course data.
type CourseRepository interface {
	// GetCourseByID retrieves a course by its ID.
	GetCourseByID(ctx context.Context, courseID string) (*dto.CourseDTO, error)

	// CreateCourse inserts a new course.
	CreateCourse(ctx context.Context, createDTO *dto.CreateCourseDTO) (*dto.CourseDTO, error)
}
