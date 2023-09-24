package service

import (
	"context"

	"github.com/moswil/course-management/internal/core/domain/event"
	"github.com/moswil/course-management/internal/core/dto"
	eventer "github.com/moswil/course-management/internal/core/interface/event"
	"github.com/moswil/course-management/internal/core/interface/repository"
	"github.com/moswil/course-management/internal/core/interface/service"
)

// CourseService is the implementation of the CourseService interface.
type CourseService struct {
	// You can inject any dependencies or repositories needed here.
	courseRepository repository.CourseRepository
	eventPublisher   eventer.EventPublisher
}

// NewCourseServicer creates a new CourseService instance.
func NewCourseService(courseRepository repository.CourseRepository, eventPublisher eventer.EventPublisher) service.CourseService {
	return &CourseService{
		courseRepository: courseRepository,
		eventPublisher:   eventPublisher,
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

	// Publish a course creation event
	createCourseEvent := event.CourseCreatedEvent{
		CourseID:   course.CourseID,
		Title:      course.Title,
		Instructor: course.Instructor,
		// Set other event fields as needed
	}

	if err := s.eventPublisher.Publish(ctx, createCourseEvent); err != nil {
		// Handle the event publishing error
		return nil, err
	}

	return course, err
}
