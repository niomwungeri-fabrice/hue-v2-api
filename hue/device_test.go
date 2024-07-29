package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	mockDevicesResponse := &DeviceList{
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
					Name:      "DeviceDetails 1",
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
					Name:      "DeviceDetails 2",
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
	if device1.Metadata.Name != "DeviceDetails 1" {
		t.Errorf("Expected device name to be 'DeviceDetails 1', got '%s'", device1.Metadata.Name)
	}
	device2 := devices.Data[1]
	if device2.ID != "2" {
		t.Errorf("Expected device ID to be '2', got '%s'", device2.ID)
	}
	if device2.ProductData.ModelID != "model2" {
		t.Errorf("Expected device model ID to be 'model2', got '%s'", device2.ProductData.ModelID)
	}
	if device2.Metadata.Name != "DeviceDetails 2" {
		t.Errorf("Expected device name to be 'DeviceDetails 2', got '%s'", device2.Metadata.Name)
	}
}

func TestGetDevice(t *testing.T) {
	// Mock response data for GetDevice
	mockDeviceResponse := &DeviceList{
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
					Name:      "DeviceDetails 1",
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
		},
	}

	// Mock HTTP client
	mockResponseBody, _ := json.Marshal(mockDeviceResponse)
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

	// Call GetDevice and check the response
	device, err := client.GetDevice("1")
	if err != nil {
		t.Fatalf("GetDevice returned an error: %v", err)
	}

	fmt.Println(device, "==========")

	if device.Data[0].ID != "1" {
		t.Errorf("Expected device ID to be '1', got '%s'", device.Data[0].ID)
	}
	if device.Data[0].ProductData.ModelID != "model1" {
		t.Errorf("Expected device model ID to be 'model1', got '%s'", device.Data[0].ProductData.ModelID)
	}
	if device.Data[0].Metadata.Name != "DeviceDetails 1" {
		t.Errorf("Expected device name to be 'DeviceDetails 1', got '%s'", device.Data[0].Metadata.Name)
	}
}

func TestUpdateDevice(t *testing.T) {
	// Mock HTTP client
	mockClient := &http.Client{
		Transport: &MockTransport{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
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

	// Create a mock update body
	updateBody := bytes.NewBufferString(`{"name": "Updated Device"}`)

	// Call UpdateDevice and check for errors
	err := client.UpdateDevice("1", updateBody)
	if err != nil {
		t.Fatalf("UpdateDevice returned an error: %v", err)
	}
}

func TestDeleteDevice(t *testing.T) {
	// Mock response data for DeleteDevice
	mockDeleteResponse := &DeleteResponse{
		Errors: nil,
	}

	// Mock HTTP client
	mockResponseBody, _ := json.Marshal(mockDeleteResponse)
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

	// Call DeleteDevice and check the response
	deleteResponse, err := client.DeleteDevice("1")
	if err != nil {
		t.Fatalf("DeleteDevice returned an error: %v", err)
	}

	if deleteResponse.Errors != nil {
		t.Errorf("Expected delete success to be true, got false")
	}
}
