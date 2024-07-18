# Hue API v2 Golang Package

This project provides a Golang package to interact `quickly` with the Philips Hue API v2. It includes functions to discover and manage Hue devices, with built-in local bridge interactions.

## Prerequisite
- Get your bridge IP: https://developers.meethue.com/develop/application-design-guidance/hue-bridge-discovery/
- Get your Username: [https://developers.meethue.com/develop/hue-api-v2/cloud2cloud-getting-started/](https://developers.meethue.com/develop/hue-api-v2/getting-started/)

## Features

- Interact with Philips Hue API v2
- Discover and manage Hue devices
- Support for both cloud and local bridge interactions
- Cross-platform builds for Linux, Windows, and macOS

## Installation

To install the package, use:

```sh
go get github.com/yourusername/hue-v2-api
```
## Available Clients

### GetDevices

Fetches the list of devices connected to the Hue bridge.

```go
package main

import (
    "fmt"
    "github.com/niomwungeri-fabrice/hue-v2-api/hue"
)

func main() {
    client, err := hue.NewClient("https://<bridge-ip>", "", "your-hue-application-key")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // true to remove bridges
    devices, err := client.GetDevices(false)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, device := range devices.Data {
        fmt.Printf("Device ID: %s, Name: %s\n", device.ID, device.Metadata.Name)
    }
}
```
## Running Tests
```sh
 go test ./hue -v
```

## Coming Soon
The package will be extended with more functionality to interact with different aspects of the Hue API v2, such as:

- Managing lights
- Controlling scenes
- Handling schedules
- And more...

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Resources:
- https://developers.meethue.com/develop/hue-api-v2/

## End Goal
Turn into the standard way of connecting to the Phillips hue v2 API. In other words production ready.

