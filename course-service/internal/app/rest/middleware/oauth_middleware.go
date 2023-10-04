package middleware

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type OAuth2Config struct {
	Issuer   string
	ClientID string
}

// claims component of jwt contains mainy fields , we need only roles of demo-client
// "resource_access":{"demo-client":{"roles":["instructor","student","course-admin"]}},
type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

type client struct {
	DemoClient clientRoles `json:"demo-client,omitempty"`
}

type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

func TokenValidationMiddleware(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")

		oauth2Config := &OAuth2Config{
			Issuer:   viper.GetString("oauth.issuer"),
			ClientID: viper.GetString("oauth.client_id"),
		}

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Timeout:   time.Duration(6000) * time.Second,
			Transport: tr,
		}

		oidcContext := oidc.ClientContext(ctx, client)
		provider, err := oidc.NewProvider(oidcContext, oauth2Config.Issuer)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization failed while getting the provider"})
			ctx.Abort()
			return
		}

		oidcConfig := &oidc.Config{
			ClientID: oauth2Config.ClientID,
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(ctx, tokenString)
		if err != nil {
			log.Printf("error verifying token: %v\n", err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorisation failed while verifying the token"})
			ctx.Abort()
			return
		}

		var IDTokenClaims Claims // ID Token payload is just JSON.
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "claims " + err.Error()})
			ctx.Abort()
			return
		}

		// checking the roles
		userAccessRoles := IDTokenClaims.ResourceAccess.DemoClient.Roles
		log.Printf("User roles: %v\n", userAccessRoles)

		roleAllowed := false
		for _, b := range userAccessRoles {
			if b == role {
				roleAllowed = true
				break
			}
		}

		if roleAllowed {
			// If the role is allowed, call the next middleware/handler.
			ctx.Next()
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "User not allowed to access this API"})
			ctx.Abort()
		}
	}
}
