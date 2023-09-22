package service

import (
	context "context"

	"github.com/moswil/course-management/internal/core/dto"
	"github.com/moswil/course-management/internal/core/interface/service"
)

// CourseService implements the gRPC CourseServiceServer interface.
type CourseService struct {
	courseService service.CourseService
}

// mustEmbedUnimplementedCourseServiceServer implements CourseServiceServer.
func (CourseService) mustEmbedUnimplementedCourseServiceServer() {
	panic("unimplemented")
}

func NewCourseService(courseService service.CourseService) *CourseService {
	return &CourseService{
		courseService: courseService,
	}
}

func (s CourseService) CreateCourse(ctx context.Context, req *CreateCourseRequest) (*CreateCourseResponse, error) {
	// let's call the core business logic to create a course
	createDTO := &dto.CreateCourseDTO{
		Title:      req.Title,
		Instructor: req.Instructor,
	}
	course, err := s.courseService.CreateCourse(ctx, createDTO)
	if err != nil {
		return nil, err
	}

	return &CreateCourseResponse{
		CourseId:   course.CourseID,
		Title:      course.Title,
		Instructor: course.Instructor,
	}, nil
}

func (s CourseService) GetCourse(ctx context.Context, req *GetCourseRequest) (*GetCourseResponse, error) {
	course, err := s.courseService.GetCourseByID(ctx, req.CourseId)
	if err != nil {
		return nil, err
	}

	return &GetCourseResponse{
		CourseId:   course.CourseID,
		Title:      course.Title,
		Instructor: course.Instructor,
	}, nil
}
