package keycloak

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
	"github.com/moswil/course-management/user-service/internal/core/dto"
	"github.com/moswil/course-management/user-service/internal/core/port/service"
)

type KeycloakProvider struct {
	gocloak      *gocloak.GoCloak
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloakProvider() service.OAuthProviderService {
	return &KeycloakProvider{
		gocloak:      gocloak.NewClient("http://dev.keycloak.shirwil.com:8180"),
		clientId:     "demo-sa-client",
		clientSecret: "1c4fe703-55f0-485c-a870-8f7480a6c1eb",
		realm:        "demo-realm",
	}
}

func (kp *KeycloakProvider) CreateUserInProvider(ctx context.Context, createUser *dto.CreateUserDTO) (string, error) {
	// login client
	token, err := kp.gocloak.LoginClient(ctx, kp.clientId, kp.clientSecret, kp.realm)
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	// Define the user attributes
	attributes := map[string][]string{
		"middleName":  {createUser.MiddleName},
		"phoneNumber": {createUser.PhoneNumber},
		// Add more custom attributes as needed
	}

	// credentials
	credential := gocloak.CredentialRepresentation{
		Type:      gocloak.StringP("password"),
		Temporary: gocloak.BoolP(false),
		Value:     gocloak.StringP(createUser.Password),
	}

	// Set the user's password using a slice containing the CredentialRepresentation
	credentials := []gocloak.CredentialRepresentation{credential}

	// user object
	user := gocloak.User{
		ID:            gocloak.StringP(uuid.New().String()),
		Username:      gocloak.StringP(createUser.Username),
		Email:         gocloak.StringP(createUser.Email),
		FirstName:     gocloak.StringP(createUser.FirstName),
		LastName:      gocloak.StringP(createUser.LastName),
		Attributes:    &attributes,
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(false),
		Credentials:   &credentials,
	}

	// TODO: attach roles to the user

	// create user
	userId, err := kp.gocloak.CreateUser(ctx, token.AccessToken, kp.realm, user)
	if err != nil {
		log.Printf("error creating user: %s\n", err.Error())
		return "", err
	}

	return userId, nil
}
