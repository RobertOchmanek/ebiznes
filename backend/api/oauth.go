package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var FrontendAddress = os.Getenv("FRONTEND_ADDRESS")
var backendAddress = os.Getenv("BACKEND_ADDRESS")
var githubClientId = os.Getenv("GITHUB_CLIENT_ID")
var githubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

func OauthConfig() *oauth2.Config {

	//Provide default configuration for oauth provider
	oauthConfig := &oauth2.Config{
		ClientID:     githubClientId,
		ClientSecret: githubClientSecret,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email", "read:user"},
		RedirectURL:  backendAddress + "/oauth/callback",
	}

	return oauthConfig
}

func OauthLoginUrl(c echo.Context) error {

	oauthUrl := OauthConfig().AuthCodeURL("state")

	return c.JSON(http.StatusOK, oauthUrl)
}

func OauthLogoutUrl(c echo.Context) error {

	//Get user ID from query param
	userId := c.Param("userId")

	//Obtain current database connection and fetch user by ID
	db := database.DbManager()

	user := model.User{}
	db.Where(database.IdEquals, userId).Find(&user)

	//Invalidate user tokens and update user in DB
	user.OauthToken = ""
	user.UserToken = ""

	db.Save(&user)

	//Redirect user to login page
	return c.JSON(http.StatusOK, FrontendAddress)
}

func OauthCallback(c echo.Context) error {

	//Request access token from provider
	oauthToken, err := OauthConfig().Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	//Create user data request
	userRequest, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userRequest.Header.Add("Accept", "application/vnd.github.v3+json")
	userRequest.Header.Add("Authorization", "token "+oauthToken.AccessToken)

	//Perform user data request
	userResponse, err := http.DefaultClient.Do(userRequest)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	defer userResponse.Body.Close()

	//Read user data from response body and cast it to string
	userData, err := ioutil.ReadAll(userResponse.Body)
	userDataString := string(userData)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	//Create temporary struct to hold user data returnerd from request
	userDataStruct := struct {
		ID    int
		Login string
	}{}

	//Convert user data json to temporary struct
	json.Unmarshal([]byte(userDataString), &userDataStruct)

	//Create new internal user token to save or refresh
	userToken := uuid.New()

	//Obtain current database connection
	db := database.DbManager()

	if !UserExists(userDataStruct.Login) {

		//Create new user if missing from DB
		newCart := model.Cart{
			CartItems: []model.CartItem{},
		}

		internalUser := model.User{}
		internalUser.Username = userDataStruct.Login
		internalUser.OauthId = userDataStruct.ID
		internalUser.OauthToken = oauthToken.AccessToken
		internalUser.UserToken = userToken.String()
		internalUser.Cart = newCart

		db.Create(&internalUser)
	} else {

		//If user exists refresh access token
		user := model.User{}
		db.Where("username = ?", userDataStruct.Login).Find(&user)

		user.OauthToken = oauthToken.AccessToken
		user.UserToken = userToken.String()

		db.Save(&user)
	}

	//Redirect the user to the home page with acces token as query param
	return c.Redirect(http.StatusFound, FrontendAddress+"?user_token="+userToken.String())
}
