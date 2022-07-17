package api

import (
	"context"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/database"
)

const clientId = ""
const clientSecret = ""

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

	//Create user data request
	userRequest, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userRequest.Header.Add("Accept", "application/vnd.github.v3+json")
	userRequest.Header.Add("Authorization", "token " + oauthToken.AccessToken)

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
		ID       int
		Login    string
	}{}

	//Convert user data json to temporary struct
	json.Unmarshal([]byte(userDataString), &userDataStruct)

	//Create new internal user token to save or refresh
	userToken := uuid.New()

	//Obtain current database connection
	db := database.DbManager()

	if (!UserExists(userDataStruct.Login)) {

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
	return c.Redirect(http.StatusFound, "http://localhost:3000?user_token=" + userToken.String())
}