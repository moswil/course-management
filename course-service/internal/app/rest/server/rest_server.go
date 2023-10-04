package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/internal/app/port/server"
	"github.com/moswil/course-management/internal/app/rest/controller"
	"github.com/moswil/course-management/internal/app/rest/middleware"
	"github.com/moswil/course-management/internal/core/port/service"
	"github.com/moswil/course-management/internal/core/util/constants"
	"github.com/moswil/course-management/internal/infra/oauth"
	"go.opentelemetry.io/otel/trace"
)

// RestServer represents the REST API server.
type RestServer struct {
	engine *gin.Engine
}

// NewRestServer creates a new RestServer instance.
func NewRestServer(courseService service.CourseService, tracer trace.Tracer) server.AppServer {
	server := &RestServer{
		engine: gin.Default(),
	}

	server.setupRoutes(courseService, tracer)
	return server
}

// setupRoutes configures the REST API routes.
func (s *RestServer) setupRoutes(courseService service.CourseService, tracer trace.Tracer) {
	courseController := controller.NewCourseController(courseService, tracer)
	// pass the oauth.NewOAuth2Config() as a parameter
	oauthController := controller.NewOAuth2Controller(oauth.NewOAuth2Config())

	// Use the middleware
	s.engine.Use(middleware.TracingMiddleware()) // adding tracing middleware

	v1 := s.engine.Group("/api/v1")
	{
		courseV1 := v1.Group("/courses")
		{
			courseV1.POST("/", middleware.TokenValidationMiddleware(constants.ROLE_INSTRUCTOR), courseController.CreateCourse)
			courseV1.GET("/:course_id", middleware.TokenValidationMiddleware(constants.ROLE_COURSE_ADMIN), courseController.GetCourse)
		}

	}

	oauth2 := s.engine.Group("/oauth2")
	{
		oauth2V1 := oauth2.Group("/")
		{
			oauth2V1.GET("/callback", oauthController.OAuth2Callback)
			oauth2V1.GET("/authorize", oauthController.OAuth2Authorize)

		}
	}
}

// Start starts the REST API server.
func (s *RestServer) Start(addr string) error {
	return s.engine.Run(addr)
}

// Stop Rest Server
func (s *RestServer) Stop() error {
	// return nil
	return errors.New("not implemented!")
}
