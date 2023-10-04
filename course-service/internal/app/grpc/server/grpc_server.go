package server

import (
	"fmt"
	"net"

	"github.com/moswil/course-management/internal/app/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer represents the gRPC server.
type GRPCServer struct {
	port          int
	courseService service.CourseService
}

// NewGRPCServer creates a new GRPCServer instance.
func NewGRPCServer(port int, courseService service.CourseService) *GRPCServer {
	return &GRPCServer{
		port:          port,
		courseService: courseService,
	}
}

// Start starts the gRPC server.
func (s *GRPCServer) Start() error {
	// create a new gRPC server
	server := grpc.NewServer()

	// Register your gRPC service implementation
	service.RegisterCourseServiceServer(server, s.courseService)

	// Register reflection service on gRPC server (for debugging and exploration)
	reflection.Register(server)

	// Create a listener on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	fmt.Printf("gRPC server is running on port %d\n", s.port)

	// Start the gRPC server
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
