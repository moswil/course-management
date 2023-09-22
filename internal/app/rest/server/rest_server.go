package server

import (
	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/internal/app/rest/controller"
	"github.com/moswil/course-management/internal/core/interface/service"
)

// RestServer represents the REST API server.
type RestServer struct {
	engine *gin.Engine
}

// NewRestServer creates a new RestServer instance.
func NewRestServer(courseService service.CourseService) *RestServer {
	server := &RestServer{
		engine: gin.Default(),
	}

	server.setupRoutes(courseService)
	return server
}

// setupRoutes configures the REST API routes.
func (s *RestServer) setupRoutes(courseService service.CourseService) {
	courseController := controller.NewCourseController(courseService)

	v1 := s.engine.Group("/api/v1")
	{
		courseV1 := v1.Group("/courses")
		{
			courseV1.POST("/", courseController.CreateCourse)
			courseV1.GET("/:course_id", courseController.GetCourse)
		}
	}
}

// Start starts the REST API server.
func (s *RestServer) Start(addr string) error {
	return s.engine.Run(addr)
}
