package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

const clientId = "<client ID>"
const clientSecret = "<client secret>"

func Authorize(c echo.Context) error {

	httpClient := http.Client{}

	//Get request token from query param
	requestToken := c.QueryParam("code")

	requestUrl := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientId, clientSecret, requestToken)

	req, err := http.NewRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	//Set accept header to obtain the response as JSON
	req.Header.Set("accept", "application/json")

	//Perform access token request
	res, err := httpClient.Do(req)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	//Parse the request body
	var accessResponse AccessResponse
	if err := json.NewDecoder(res.Body).Decode(&accessResponse); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	//Redirect the user to the home page with acces token as query param
	return c.Redirect(http.StatusFound, "http://localhost:3000?access_token=" + accessResponse.AccessToken)
}

type AccessResponse struct {
	AccessToken string `json:"access_token"`
}