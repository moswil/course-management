package model

import (
	"github.com/moswil/course-management/internal/core/domain/entity"
	"github.com/moswil/course-management/internal/core/dto"
)

// TODO: Use an adapter pattern for `CourseModel` from `model.Course` entity.

// CourseModel represents the data model for storing courses in the database using GORM.
type CourseModel struct {
	CourseID   string `gorm:"primaryKey;column:course_id"`
	Title      string `gorm:"column:title"`
	Instructor string
	// Add more fields as needed
}

// TableName sets the table name for the CourseModel in the database.
func (CourseModel) TableName() string {
	return "courses" // Customize the table name as needed
}

// ToEntity converts a CourseModel to a Course entity
func (cm *CourseModel) ToEntity() *entity.Course {
	courseEntity := &entity.Course{}
	courseEntity.SetCourseID(cm.CourseID)
	courseEntity.SetTitle(cm.Title)
	courseEntity.SetInstructor(cm.Instructor)
	return courseEntity
}

// FromEntity converts a CourseEntity to a CourseModel.
func (cm *CourseModel) FromEntity(courseEntity *entity.Course) {
	cm.CourseID = courseEntity.GetCourseID()
	cm.Title = courseEntity.GetTitle()
	cm.Instructor = courseEntity.GetInstructor()
}

// ToDTO converts a CourseModel to a Course DTO
func (cm *CourseModel) ToDTO() *dto.CourseDTO {
	return &dto.CourseDTO{
		CourseID:   cm.CourseID,
		Title:      cm.Title,
		Instructor: cm.Instructor,
	}
}
