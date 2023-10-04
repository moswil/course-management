package event

import "time"

type CourseEnrollmentEvent struct {
	CourseID       string
	StudentID      string
	EnrollmentTime time.Time
}

func NewCourseEnrollmentEvent(courseID, studentID string) *CourseEnrollmentEvent {
	return &CourseEnrollmentEvent{
		CourseID:       courseID,
		StudentID:      studentID,
		EnrollmentTime: time.Now(),
	}
}
