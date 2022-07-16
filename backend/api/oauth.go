package api

import (
	"context"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"github.com/labstack/echo/v4"
)

const clientId = "clientId"
const clientSecret = "clientSecret"

func OauthConfig() *oauth2.Config {

	//Provide default configuration for oauth provider
	oauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email","read:user",},
		RedirectURL:  "http://localhost:8080/oauth/callback",
	}

	return oauthConfig
}

func OauthUrl(c echo.Context) error {

	oauthUrl := OauthConfig().AuthCodeURL("state")

	return c.JSON(http.StatusOK, oauthUrl)
}

func OauthCallback(c echo.Context) error {

	//Request access token from provider
	oauthToken, err := OauthConfig().Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	
	//Redirect the user to the home page with acces token as query param
	return c.Redirect(http.StatusFound, "http://localhost:3000?access_token=" + oauthToken.AccessToken)
}