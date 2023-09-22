package service

import (
	"context"

	"github.com/moswil/course-management/internal/core/dto"
	"github.com/moswil/course-management/internal/core/interface/repository"
	"github.com/moswil/course-management/internal/core/interface/service"
)

// CourseService is the implementation of the CourseService interface.
type CourseService struct {
	// You can inject any dependencies or repositories needed here.
	courseRepository repository.CourseRepository
}

// NewCourseServicer creates a new CourseService instance.
func NewCourseService(courseRepository repository.CourseRepository) service.CourseService {
	return &CourseService{
		courseRepository: courseRepository,
	}
}

// GetCourseByID retrieves a course by its ID.
func (s *CourseService) GetCourseByID(ctx context.Context, courseID string) (*dto.CourseDTO, error) {
	// Use the courseRepository to fetch the course data.
	course, err := s.courseRepository.GetCourseByID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

// CreateCourse creates a new course using a CreateCourseDTO.
func (s *CourseService) CreateCourse(ctx context.Context, createDTO *dto.CreateCourseDTO) (*dto.CourseDTO, error) {
	// Use the courseRepository to create a new course in the database.
	course, err := s.courseRepository.CreateCourse(ctx, createDTO)
	if err != nil {
		// do some extra logic
		return nil, err
	}
	return course, err
}
