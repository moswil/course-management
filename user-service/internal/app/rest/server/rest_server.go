package server

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/moswil/course-management/internal/app/port/server"
	"github.com/moswil/course-management/user-service/internal/app/rest/controller"
	"github.com/moswil/course-management/user-service/internal/core/port/service"
)

// RestServer represents the REST API server.
type RestServer struct {
	engine *gin.Engine
}

// NewRestServer creates a new RestServer instance.
func NewRestServer(userService service.UserService) server.AppServer {
	server := &RestServer{
		engine: gin.Default(),
	}

	server.setupRoutes(userService)
	return server
}

// setupRoutes configures the REST API routes.
func (s *RestServer) setupRoutes(userService service.UserService) {
	courseController := controller.NewUserController(userService)

	v1 := s.engine.Group("/api/v1")
	{
		courseV1 := v1.Group("/users")
		{
			courseV1.POST("/", courseController.CreateUser)
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
