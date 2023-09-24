package event

import "fmt"

type EventHandler interface {
	HandleCourseCreatedEvent(event CourseCreatedEvent) error
	// Add handlers for other events as needed
}

type CourseEventHandler struct {
	// Implement your event handling logic here
}

func NewCourseEventHandler() *CourseEventHandler {
	// Initialize your event handler
	return &CourseEventHandler{}
}

func (h *CourseEventHandler) HandleCourseCreatedEvent(event CourseCreatedEvent) error {
	// Implement the logic to handle the course creation event
	fmt.Printf("Course Created Event Received: ID=%s, Title=%s, Instructor=%s\n", event.CourseID, event.Title, event.Instructor)
	// For example, update a read model or perform other actions
	return nil
}
