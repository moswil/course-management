package service

import (
	"context"

	"github.com/moswil/course-management/internal/core/dto"
)

// CourseService defines the methods for managing courses.
type CourseService interface {
	// GetCourseByID retrieves a course by its ID.
	GetCourseByID(ctx context.Context, courseID string) (*dto.CourseDTO, error)

	// CreateCourse creates a new course.
	CreateCourse(ctx context.Context, course *dto.CreateCourseDTO) (*dto.CourseDTO, error)
}
