package hue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Device struct {
	ID          string `json:"id"`
	IDV1        string `json:"id_v1,omitempty"`
	ProductData struct {
		ModelID              string  `json:"model_id"`
		ManufacturerName     string  `json:"manufacturer_name"`
		ProductName          string  `json:"product_name"`
		ProductArchetype     string  `json:"product_archetype"`
		Certified            bool    `json:"certified"`
		SoftwareVersion      string  `json:"software_version"`
		HardwarePlatformType *string `json:"hardware_platform_type,omitempty"`
	} `json:"product_data"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
	} `json:"metadata"`
	Identify struct{} `json:"identify"`
	Services []struct {
		RID   string `json:"rid"`
		RType string `json:"rtype"`
	} `json:"services"`
	Type string `json:"type"`
}

// DeviceList list of devices available for the user
type DeviceList struct {
	Errors []interface{} `json:"errors"`
	Data   []Device      `json:"data"`
}

// DeviceDetails details of one device
type DeviceDetails struct {
	Errors []interface{} `json:"errors"`
	Data   []Device      `json:"data"`
}

// GetDevices get list of devices
func (c *Client) GetDevices(removeBridge bool) (*DeviceList, error) {
	req, err := c.newRequest(http.MethodGet, "/clip/v2/resource/device", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get devices")
	}

	var devices *DeviceList
	err = json.NewDecoder(res.Body).Decode(&devices)
	if err != nil {
		return nil, err
	}
	// TODO: work on how to filter out bridge
	return devices, nil
}

// GetDevice get device details
func (c *Client) GetDevice(deviceID string) (*DeviceDetails, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/clip/v2/resource/device/%s", deviceID), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get device")
	}
	var device *DeviceDetails
	err = json.NewDecoder(res.Body).Decode(&device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

// DeleteDevice delete a device
func (c *Client) DeleteDevice(deviceID string) (*DeleteResponse, error) {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/clip/v2/resource/device/%s", deviceID), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get device")
	}

	var deleteResponse *DeleteResponse
	err = json.NewDecoder(res.Body).Decode(&deleteResponse)
	if err != nil {
		return nil, err
	}
	return deleteResponse, nil
}

// UpdateDevice update device
func (c *Client) UpdateDevice(deviceID string, updateBody io.Reader) error {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/clip/v2/resource/device/%s", deviceID), updateBody)
	if err != nil {
		return err
	}

	fmt.Println(req)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to update device")
	}
	fmt.Println("Successfully updated device")
	return nil
}
