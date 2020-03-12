package client

import (
	"bytes"
	"encoding/json"
	"errors"
)

// PathLogin is e-mas login endpoint.
const PathLogin = "/v1/sso/login"

// LoginRequest is model to request login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// LoginResponse is response model for login.
type LoginResponse struct {
	BaseResponse
	Data *struct {
		CustomerData
		Token  string `json:"token"`
		Avatar string `json:"avatar"`
	} `json:"data"`
}

// CustomerData is customer data returned after login.
// Can be used by clients to provide general customer
// data for registering in their website.
type CustomerData struct {
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Name   string `json:"name"`
}

// Login to login e-mas SSO with username/password or token.
func (c *Client) Login(userPass ...string) (data CustomerData, err error) {
	// E-mas login endpoint.
	url := c.DomainURL + PathLogin

	// Prepare request data.
	var loginRequest LoginRequest

	// Use access token if exists.
	if c.Token != "" {
		loginRequest.Token = c.Token
	}

	// Use username & password if provided.
	if len(userPass) == 2 {
		loginRequest.Username = userPass[0]
		loginRequest.Password = userPass[1]
	}

	// Required login request data not found.
	if loginRequest.Username == "" && loginRequest.Password == "" && loginRequest.Token == "" {
		return data, errors.New("required username/password or token")
	}

	// Prepare request.
	reqBody, _ := json.Marshal(loginRequest)
	request, err := c.NewRequest("POST", url, bytes.NewBuffer(reqBody), c.GetHeader())
	if err != nil {
		return data, err
	}

	// Execute the request.
	_, body, err := c.ExecuteRequest(request)
	if err != nil {
		return data, err
	}

	// Convert response to struct.
	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return data, err
	}

	// If not success.
	if loginResponse.Message != SuccessMessage {
		return data, errors.New(loginResponse.Message)
	}

	// Keep the token in client.
	c.Token = loginResponse.Data.Token

	// Prepare customer data return.
	data.Email = loginResponse.Data.Email
	data.Phone = loginResponse.Data.Phone
	data.Name = loginResponse.Data.Name

	return data, nil
}
