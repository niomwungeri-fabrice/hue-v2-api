package hue

import (
    "encoding/json"
    "errors"
    "net/http"
)

type Device struct {
    Id          string `json:"id"`
    Metadata    Metadata `json:"metadata"`
    ProductData ProductData `json:"product_data"`
}

type Metadata struct {
    Name string `json:"name"`
}

type ProductData struct {
    Manufacturer string `json:"manufacturer"`
    ModelId      string `json:"model_id"`
}

func (c *Client) GetDevices() ([]Device, error) {
    req, err := c.newRequest("GET", "/resource/device")
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

    var devices []Device
    err = json.NewDecoder(res.Body).Decode(&devices)
    if err != nil {
        return nil, err
    }

    return devices, nil
}
