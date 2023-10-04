package request

import "github.com/moswil/course-management/internal/core/dto"

// CreateCourseRequest request object for creating a course.
type CreateCourseRequest struct {
	Course dto.CreateCourseDTO `json:"course"`
}

// GetCourseRequest request object for retrieving a course.
type GetCourseRequest struct {
	CourseID string `json:"course_id"`
}
