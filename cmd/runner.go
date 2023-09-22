package main

import (
	"fmt"
	"log"

	grpc "github.com/moswil/course-management/internal/app/grpc/server"
	grpcSvc "github.com/moswil/course-management/internal/app/grpc/service"
	rest "github.com/moswil/course-management/internal/app/rest/server"
	repo "github.com/moswil/course-management/internal/core/interface/repository"
	"github.com/moswil/course-management/internal/core/service"
	"github.com/moswil/course-management/internal/infra/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize repositories, services, and servers
	courseRepository, err := initializeMySQLRepository()
	if err != nil {
		log.Printf("error: %v occurred connecting to db", err)
	}
	courseService := service.NewCourseService(courseRepository)

	// Start gRPC Server as a goroutine
	go func() {
		grpcPort := 50052
		grpcServer := grpc.NewGRPCServer(grpcPort, *grpcSvc.NewCourseService(courseService))
		if err := grpcServer.Start(); err != nil {
			fmt.Printf("Failed to start gRPC server: %v\n", err)
		}
	}()

	// Start the REST API server
	go func() {
		restServer := rest.NewRestServer(courseService)
		if err := restServer.Start(":8080"); err != nil {
			log.Fatalf("Failed to start REST server: %v", err)
		}
	}()

	// optionally wait for the go routines
	select {}
}

func initializeMySQLRepository() (repo.CourseRepository, error) {
	// Replace with your actual MySQL database connection details
	dsn := "codeguru:CodeGuru96@tcp(localhost:3306)/courses"

	// Initialize GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the Course model
	if err := db.AutoMigrate(&repository.CourseModel{}); err != nil {
		return nil, err
	}

	// Initialize the GORM repository
	courseRepository := repository.NewMySQLCourseRepository(db)

	return courseRepository, nil
}
