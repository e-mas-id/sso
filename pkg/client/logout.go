package client

import (
	"encoding/json"
	"errors"
)

// PathLogout is e-mas loogout endpoint.
const PathLogout = "/v1/sso/logout"

// Logout to logout customer SSO login. Will invalidate current token.
func (c *Client) Logout() (err error) {
	// E-mas logout endpoint.
	url := c.DomainURL + PathLogout

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
	var logoutResponse BaseResponse
	err = json.Unmarshal(body, &logoutResponse)
	if err != nil {
		return err
	}

	// If not success.
	if logoutResponse.Code != SuccessCode {
		return errors.New(logoutResponse.Message)
	}

	// Remove token.
	c.Token = ""

	return nil
}
