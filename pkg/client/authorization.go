package client

import (
	"encoding/json"
	"errors"
)

// PathAuthorization is e-mas authorization endpoint.
const PathAuthorization = "/v1/sso/auth"

// Authorize to authoorize customer SSO login for the first time.
func (c *Client) Authorize() (err error) {
	// E-mas authorize endpoint.
	url := c.DomainURL + PathAuthorization

	// Token not found.
	if c.Token == "" {
		return errors.New("required token")
	}

	// Prepare request.
	request, err := c.NewRequest("GET", url, nil, c.GetHeader())
	if err != nil {
		return err
	}

	// Add token in query param.
	c.AddQueryParam(request, "token", c.Token)

	// Execute the request.
	_, body, err := c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	// Convert response to struct.
	var authResponse BaseResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return err
	}

	// If not success.
	if authResponse.Code != SuccessCode {
		return errors.New(authResponse.Message)
	}

	return nil
}
