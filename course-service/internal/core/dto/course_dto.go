package dto

// CreateCourseDTO provides a DTO for creating a course.
type CreateCourseDTO struct {
	Title      string `json:"title"`
	Instructor string `json:"instructor"`
}

// CourseDTO for retrieving a course
type CourseDTO struct {
	CourseID   string `json:"course_id"`
	Title      string `json:"title"`
	Instructor string `json:"instructor"`
}
