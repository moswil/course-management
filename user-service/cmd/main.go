package main

import (
	"log"

	rest "github.com/moswil/course-management/user-service/internal/app/rest/server"
	"github.com/moswil/course-management/user-service/internal/core/service"
	"github.com/moswil/course-management/user-service/internal/infra/oauth_provider/keycloak"
	"github.com/moswil/course-management/user-service/internal/infra/util"
)

func main() {
	// Initialize the database to use
	// we're using MySQL
	userRepository, err := util.InitializeMySQLRepository()
	if err != nil {
		log.Printf("error: %v occurred connecting to db", err)
	}

	// choose your OAuth Provider
	oauthService := keycloak.NewKeycloakProvider()

	// inject the dependencies to the user service
	userService := service.NewUserService(userRepository, oauthService)

	// controller (rest controller)
	restServer := rest.NewRestServer(userService)
	if err := restServer.Start(":8081"); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
