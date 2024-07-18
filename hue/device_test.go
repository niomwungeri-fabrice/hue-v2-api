package hue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

// MockTransport is a custom RoundTripper for mocking HTTP responses
type MockTransport struct {
	Response *http.Response
	Error    error
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.Response, m.Error
}

func TestGetDevices(t *testing.T) {
	// Mock response data
	mockDevicesResponse := &Devices{
		Errors: []interface{}{},
		Data: []Device{
			{
				ID: "1",
				ProductData: struct {
					ModelID              string  `json:"model_id"`
					ManufacturerName     string  `json:"manufacturer_name"`
					ProductName          string  `json:"product_name"`
					ProductArchetype     string  `json:"product_archetype"`
					Certified            bool    `json:"certified"`
					SoftwareVersion      string  `json:"software_version"`
					HardwarePlatformType *string `json:"hardware_platform_type,omitempty"`
				}{
					ModelID:          "model1",
					ManufacturerName: "manufacturer1",
					ProductName:      "product1",
					ProductArchetype: "archetype1",
					Certified:        true,
					SoftwareVersion:  "1.0.0",
				},
				Metadata: struct {
					Name      string `json:"name"`
					Archetype string `json:"archetype"`
				}{
					Name:      "Device 1",
					Archetype: "archetype1",
				},
				Services: []struct {
					RID   string `json:"rid"`
					RType string `json:"rtype"`
				}{
					{RID: "service1", RType: "type1"},
				},
				Type: "light",
			},
			{
				ID: "2",
				ProductData: struct {
					ModelID              string  `json:"model_id"`
					ManufacturerName     string  `json:"manufacturer_name"`
					ProductName          string  `json:"product_name"`
					ProductArchetype     string  `json:"product_archetype"`
					Certified            bool    `json:"certified"`
					SoftwareVersion      string  `json:"software_version"`
					HardwarePlatformType *string `json:"hardware_platform_type,omitempty"`
				}{
					ModelID:          "model2",
					ManufacturerName: "manufacturer2",
					ProductName:      "product2",
					ProductArchetype: "archetype2",
					Certified:        true,
					SoftwareVersion:  "1.1.0",
				},
				Metadata: struct {
					Name      string `json:"name"`
					Archetype string `json:"archetype"`
				}{
					Name:      "Device 2",
					Archetype: "archetype2",
				},
				Services: []struct {
					RID   string `json:"rid"`
					RType string `json:"rtype"`
				}{
					{RID: "service2", RType: "type2"},
				},
				Type: "light",
			},
		},
	}

	// Mock HTTP client
	mockResponseBody, _ := json.Marshal(mockDevicesResponse)
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

	// Call GetDevices and check the response
	devices, err := client.GetDevices(false)
	if err != nil {
		t.Fatalf("GetDevices returned an error: %v", err)
	}

	if len(devices.Data) != 2 {
		t.Errorf("Expected 2 devices, got %d", len(devices.Data))
	}

	device1 := devices.Data[0]
	if device1.ID != "1" {
		t.Errorf("Expected device ID to be '1', got '%s'", device1.ID)
	}
	if device1.ProductData.ModelID != "model1" {
		t.Errorf("Expected device model ID to be 'model1', got '%s'", device1.ProductData.ModelID)
	}
	if device1.Metadata.Name != "Device 1" {
		t.Errorf("Expected device name to be 'Device 1', got '%s'", device1.Metadata.Name)
	}
	device2 := devices.Data[1]
	if device2.ID != "2" {
		t.Errorf("Expected device ID to be '2', got '%s'", device2.ID)
	}
	if device2.ProductData.ModelID != "model2" {
		t.Errorf("Expected device model ID to be 'model2', got '%s'", device2.ProductData.ModelID)
	}
	if device2.Metadata.Name != "Device 2" {
		t.Errorf("Expected device name to be 'Device 2', got '%s'", device2.Metadata.Name)
	}
}
