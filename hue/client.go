package hue

import (
	"errors"
	"net/http"
	"regexp"
)


// Client is the structure for making requests to the Hue v2 API.
type Client struct {
    httpClient       *http.Client
    baseURL          string
    bearerToken      string
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

	// If bearerToken is not provided, log a warning (or handle it as needed)
	if bearerToken == "" && !isLocalBridge(baseURL) {
		return nil, errors.New("bearerToken is required for cloud API")
	}

	return &Client{
		httpClient:        &http.Client{},
		baseURL:           baseURL,
		bearerToken:       bearerToken,
		hueApplicationKey: hueApplicationKey,
	}, nil
}

// isLocalBridge checks if the baseURL starts with a digit value, indicating a local bridge IP address.
func isLocalBridge(baseURL string) bool {
	matched, _ := regexp.MatchString(`^\d`, baseURL)
	return matched
}

func (c *Client) newRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+url, nil)
	if err != nil {
		return nil, err
	}
	// Add headers only if it's not a local bridge
	if !isLocalBridge(c.baseURL) {
		req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
	req.Header.Add("hue-application-key", c.hueApplicationKey)
	return req, nil
}
