package oauth

import "golang.org/x/oauth2"

func NewOAuth2Config() *oauth2.Config {
	// TODO: externalize the configs
	return &oauth2.Config{
		// authorization endpoint
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://dev.keycloak.shirwil.com:8180/realms/demo-realm/protocol/openid-connect/auth",
			TokenURL: "http://dev.keycloak.shirwil.com:8180/realms/demo-realm/protocol/openid-connect/token",
		},
		// client id and secret
		ClientID:     "demo-client",
		ClientSecret: "1c4fe703-55f0-485c-a870-8f7480a6c1eb",
		// redirect url
		RedirectURL: "http://localhost:8080/oauth2/callback",
		// scopes required by the application
		// Scopes: []string{"read", "write"},
	}
}
