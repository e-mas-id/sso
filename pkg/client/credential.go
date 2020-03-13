package client

import (
	"bytes"
	"encoding/json"
	"errors"
)

// PathCredentialRequest is e-mas credential request endpoint.
const PathCredentialRequest = "/v1/sso/credential"

// CredentialRequest is model for requesting client credentials.
type CredentialRequest struct {
	ClientName string `json:"client_name"`
}

// CredentialResponse is response model for credential request.
type CredentialResponse struct {
	BaseResponse
	Data *struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"data"`
}

// RequestCredential to request client credential (client_id & client_secret) to e-mas.
func (c *Client) RequestCredential() (client *Client, err error) {
	// E-mas credential request endpoint.
	url := c.DomainURL + PathCredentialRequest

	// Credential request data.
	reqBody, _ := json.Marshal(CredentialRequest{
		ClientName: c.ClientName,
	})

	// Prepare request.
	request, err := c.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return c, err
	}

	// Execute the request.
	resp, body, err := c.ExecuteRequest(request)
	if err != nil {
		return c, err
	}

	// Response not success.
	if resp.StatusCode != SuccessCode {
		return c, errors.New(resp.Status)
	}

	// Convert response to struct.
	var credResponse CredentialResponse
	err = json.Unmarshal(body, &credResponse)
	if err != nil {
		return c, err
	}

	// Update client data.
	c.ClientId = credResponse.Data.ClientId
	c.ClientSecret = credResponse.Data.ClientSecret

	return c, nil
}
