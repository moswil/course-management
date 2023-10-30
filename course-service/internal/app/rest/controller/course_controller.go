package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/internal/app/rest/model/request"
	"github.com/moswil/course-management/internal/app/rest/model/response"
	"github.com/moswil/course-management/internal/core/port/service"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var log = otelzap.L()

// CourseController handles HTTP requests related to courses.
type CourseController struct {
	courseService service.CourseService
	tracer        trace.Tracer
}

// NewCourseController creates a new CourseController instance.
func NewCourseController(courseService service.CourseService, tracer trace.Tracer) *CourseController {
	return &CourseController{
		courseService: courseService,
		tracer:        tracer,
	}
}

// CreateCourse handles the creation of a new course.
func (c *CourseController) CreateCourse(ctx *gin.Context) {
	// add telemetry
	_ctx := ctx.Request.Context()

	// Start a custom span for the "CreateCourse" operation
	_ctx, span := c.tracer.Start(_ctx, "create course endpoint")
	ctx.Request.WithContext(_ctx)
	defer span.End()

	var createRequest request.CreateCourseRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createDTO := createRequest.Course

	course, err := c.courseService.CreateCourse(ctx, &createDTO)
	log.InfoContext(ctx, "error: "+err.Error())
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
	// add telemetry
	_ctx := ctx.Request.Context()
	// log := otelzap.L()

	// Start a custom span for the "GetCourse" operation
	_ctx, span := c.tracer.Start(_ctx, "get course endpoint")
	ctx.Request.WithContext(_ctx)
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	courseID := ctx.Param("course_id")

	// Call the service to retrieve the course by ID
	course, err := c.courseService.GetCourseByID(ctx, courseID)
	if err != nil {
		log.InfoContext(ctx, "error getting course: "+err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	res := &response.GetCourseResponse{
		Course: *course,
	}

	ctx.JSON(http.StatusOK, res)
}
