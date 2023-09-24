package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/internal/app/rest/model/request"
	"github.com/moswil/course-management/internal/app/rest/model/response"
	"github.com/moswil/course-management/internal/core/interface/service"
)

// CourseController handles HTTP requests related to courses.
type CourseController struct {
	courseService service.CourseService
}

// NewCourseController creates a new CourseController instance.
func NewCourseController(courseService service.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

// CreateCourse handles the creation of a new course.
func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var createRequest request.CreateCourseRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createDTO := createRequest.Course

	course, err := c.courseService.CreateCourse(ctx, &createDTO)
	log.Printf("error: %v", err)
	// Call the service to create the course
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course", "message": err.Error()})
		return
	}

	res := &response.CreateCourseResponse{
		Course: *course,
	}

	ctx.JSON(http.StatusCreated, res)
}

// GetCourse handles the retrieval of a course by ID.
func (c *CourseController) GetCourse(ctx *gin.Context) {
	courseID := ctx.Param("course_id")

	// Call the service to retrieve the course by ID
	course, err := c.courseService.GetCourseByID(ctx, courseID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	res := &response.GetCourseResponse{
		Course: *course,
	}

	ctx.JSON(http.StatusOK, res)
}
