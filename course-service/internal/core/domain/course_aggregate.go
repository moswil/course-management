package domain

import "github.com/moswil/course-management/internal/core/domain/entity"

type CourseAggregate struct {
	Course      *entity.Course
	Enrollments []*entity.Course
}
