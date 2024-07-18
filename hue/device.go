package hue

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HueDevices response from discovering hue's devices
type Devices struct {
	Errors []interface{} `json:"errors"`
	Data   []Device      `json:"data"`
}
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

func (c *Client) GetDevices(removeBridge bool) (*Devices, error) {
	req, err := c.newRequest("GET", "/clip/v2/resource/device")
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

	var devices *Devices
	err = json.NewDecoder(res.Body).Decode(&devices)
	if err != nil {
		return nil, err
	}
	// TODO: work on how to filter out bridge
	return devices, nil
}
