package hue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetLights(t *testing.T) {
	// Mock response data
	mockLightsResponse := &Lights{
		Errors: []interface{}{},
		Data: []Light{
			{
				ID: "1",
				ProductData: struct {
					Function string `json:"function"`
				}{
					Function: "functional1",
				},
				Metadata: struct {
					Name      string `json:"name"`
					Archetype string `json:"archetype"`
					Function  string `json:"function"`
				}{
					Name:      "Light 1",
					Archetype: "archetype1",
					Function:  "function1",
				},
				On: struct {
					On bool `json:"on"`
				}{
					On: true,
				},
				Dimming: struct {
					Brightness  float64 `json:"brightness"`
					MinDimLevel float64 `json:"min_dim_level"`
				}{
					Brightness:  100.0,
					MinDimLevel: 10.0,
				},
				Type: "light",
			},
			{
				ID: "2",
				ProductData: struct {
					Function string `json:"function"`
				}{
					Function: "functional2",
				},
				Metadata: struct {
					Name      string `json:"name"`
					Archetype string `json:"archetype"`
					Function  string `json:"function"`
				}{
					Name:      "Light 2",
					Archetype: "archetype2",
					Function:  "function2",
				},
				On: struct {
					On bool `json:"on"`
				}{
					On: false,
				},
				Dimming: struct {
					Brightness  float64 `json:"brightness"`
					MinDimLevel float64 `json:"min_dim_level"`
				}{
					Brightness:  50.0,
					MinDimLevel: 5.0,
				},
				Type: "light",
			},
		},
	}

	// Mock HTTP client
	mockResponseBody, _ := json.Marshal(mockLightsResponse)
	mockClient := &http.Client{
		Transport: &MockTransport{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(mockResponseBody)),
				Header:     make(http.Header),
			},
			Error: nil,
		},
	}

	// Create a new client with the mock HTTP client
	client := &Client{
		httpClient:        mockClient,
		baseURL:           "https://mocked-url",
		bearerToken:       "mocked-token",
		hueApplicationKey: "mocked-key",
	}

	// Call GetLights and check the response
	lights, err := client.GetLights()
	if err != nil {
		t.Fatalf("GetLights returned an error: %v", err)
	}

	if len(lights.Data) != 2 {
		t.Errorf("Expected 2 lights, got %d", len(lights.Data))
	}

	light1 := lights.Data[0]
	if light1.ID != "1" {
		t.Errorf("Expected light ID to be '1', got '%s'", light1.ID)
	}
	if light1.ProductData.Function != "functional1" {
		t.Errorf("Expected light function to be 'functional1', got '%s'", light1.ProductData.Function)
	}
	if light1.Metadata.Name != "Light 1" {
		t.Errorf("Expected light name to be 'Light 1', got '%s'", light1.Metadata.Name)
	}
	light2 := lights.Data[1]
	if light2.ID != "2" {
		t.Errorf("Expected light ID to be '2', got '%s'", light2.ID)
	}
	if light2.ProductData.Function != "functional2" {
		t.Errorf("Expected light function to be 'functional2', got '%s'", light2.ProductData.Function)
	}
	if light2.Metadata.Name != "Light 2" {
		t.Errorf("Expected light name to be 'Light 2', got '%s'", light2.Metadata.Name)
	}
}
