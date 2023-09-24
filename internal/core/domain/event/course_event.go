package event

type CourseCreatedEvent struct {
	CourseID   string `json:"course_id"`
	Title      string `json:"title"`
	Instructor string `json:"instructor"`
	// Add more fields as needed
}
