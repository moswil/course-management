package domain

import "github.com/google/uuid"

// Course entity
type Course struct {
	courseID   string
	title      string
	instructor string
}

// NewCourse creates a new course with a generated UUID as the identifier.
func NewCourse(title, instructor string) *Course {
	return &Course{
		courseID:   uuid.New().String(),
		title:      title,
		instructor: instructor,
	}
}

// GetCourseID returns a course ID
func (c *Course) GetCourseID() string {
	return c.courseID
}

// GetTitle returns the course title
func (c *Course) GetTitle() string {
	return c.title
}

// GetInstructor returns the instructor of the course
func (c *Course) GetInstructor() string {
	return c.instructor
}

// GetCourseID returns a course ID
func (c *Course) SetCourseID(courseID string) {
	c.courseID = courseID
}

// GetTitle returns the course title
func (c *Course) SetTitle(title string) {
	c.title = title
}

// GetInstructor returns the instructor of the course
func (c *Course) SetInstructor(instructor string) {
	c.instructor = instructor
}
