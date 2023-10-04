package response

import "github.com/moswil/course-management/internal/core/dto"

// CreateCourseResponse response object for creating a course.
type CreateCourseResponse struct {
	Course dto.CourseDTO `json:"course"`
}

// GetCourseResponse response object for retrieving a course.
type GetCourseResponse struct {
	Course dto.CourseDTO `json:"course"`
}
