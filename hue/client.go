package hue

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Client is the structure for making requests to the Hue v2 API.
type Client struct {
	httpClient        *http.Client
	baseURL           string
	bearerToken       string
	hueApplicationKey string
}

// NewClient creates a new Hue API client.
// If the baseURL is a local bridge IP address, you don't need to provide a bearer token.
// This is useful for local testing.
func NewClient(baseURL, bearerToken, hueApplicationKey string) (*Client, error) {
	// Validate the inputs
	if baseURL == "" {
		return nil, errors.New("baseURL is required")
	}
	if hueApplicationKey == "" {
		return nil, errors.New("hueApplicationKey is required")
	}

	if !isValidBaseURL(baseURL) {
		fmt.Println("warning: invalid base url, example base-url: https://10.0.0.250")
		baseURL = "https://api.meethue.com"
	}

	// If bearerToken is not provided, log a warning (or handle it as needed)
	if bearerToken == "" && !isValidBaseURL(baseURL) {
		return nil, errors.New("bearerToken is required for cloud API")
	} else {
		// Disable security check for local testing
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &Client{
		httpClient:        &http.Client{},
		baseURL:           baseURL,
		bearerToken:       bearerToken,
		hueApplicationKey: hueApplicationKey,
	}, nil
}

func (c *Client) newRequest(method, url string, updateRequest io.Reader) (*http.Request, error) {
	fullURL := c.baseURL + url
	req, err := http.NewRequest(method, fullURL, updateRequest)
	if err != nil {
		return nil, err
	}
	// Add headers only if it's not a local bridge
	if !isValidBaseURL(c.baseURL) {
		req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
	req.Header.Add("hue-application-key", c.hueApplicationKey)
	return req, nil
}
