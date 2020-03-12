package client

import (
	"encoding/base64"
	"github.com/e-mas-id/sso/pkg/utils"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// SuccessCode is success response header code.
	SuccessCode = 200
	// SuccessMessage is success response message.
	SuccessMessage = "success"
	// InternalErrorCode is internal response header code.
	InternalErrorCode = 500
)

// BaseResponse is base response format from e-mas SSO.
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetHeader to create header for authorization.
func (c *Client) GetHeader() (header map[string]string) {
	header = make(map[string]string)
	header["Client-Name"] = c.ClientName
	header["Authorization"] = c.getAuthHeader()
	return header
}

// getAuthHeader to encode client credential for header authorization.
func (c *Client) getAuthHeader() string {
	auth := base64.StdEncoding.EncodeToString([]byte(c.ClientId + ":" + c.ClientSecret))
	return "Basic " + auth
}

// NewRequest to create new request for SSO.
func (c *Client) NewRequest(method string, path string, body io.Reader, headers ...map[string]string) (*http.Request, error) {
	// Init http new request.
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	// Set default header.
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// Add additional headers if provided.
	if len(headers) > 0 {
		for hKey, hValue := range headers[0] {
			req.Header.Add(hKey, hValue)
		}
	}

	return req, nil
}

// ExecuteRequest to start http request.
func (c *Client) ExecuteRequest(req *http.Request) (resp *http.Response, body []byte, err error) {
	startTime := time.Now()

	// Init http client.
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Just do it.
	resp, err = httpClient.Do(req)
	utils.Log(utils.DebugType, utils.LogFmt(req.Method, time.Since(startTime).String(), req.URL.Host+req.URL.Path), c.Debug)
	if err != nil {
		return resp, nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err = ioutil.ReadAll(resp.Body)

	// Let's utils the body response.
	utils.Log(utils.DebugType, string(body), c.Debug)

	return resp, body, err
}
