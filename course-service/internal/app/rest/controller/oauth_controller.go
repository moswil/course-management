package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type OAuth2Controller struct {
	oauth2Config *oauth2.Config
}

func NewOAuth2Controller(oauth2Config *oauth2.Config) *OAuth2Controller {
	return &OAuth2Controller{
		oauth2Config: oauth2Config,
	}
}

// OAuth2Authorize handles the OAuth 2.1 authorization request
func (c *OAuth2Controller) OAuth2Authorize(ctx *gin.Context) {
	// Generate the authorization URL and redirect the user to it
	authURL := c.oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusFound, authURL)
}

func (c *OAuth2Controller) OAuth2Callback(ctx *gin.Context) {
	// Parse the authorization code from the query string
	code := ctx.DefaultQuery("code", "")

	// Exchange the authorization code for an access token
	token, err := c.oauth2Config.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// You can now use the `token` to make authenticated API requests
	// or store it securely for later use.

	// For example, to make an API request using the token:
	// client := oauth2Config.Client(c, token)
	// resp, err := client.Get("https://api.example.com/resource")
	// ...

	ctx.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken})
}
