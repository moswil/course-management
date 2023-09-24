package valueobject

type CourseID struct {
	code string
	name string
}

func NewCourseID(code, name string) *CourseID {
	return &CourseID{
		code: code,
		name: name,
	}
}
