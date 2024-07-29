# hue-v2-api

[![MIT License](https://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](https://opensource.org/licenses/MIT)

This project provides a Golang package to interact `quickly` with the Philips Hue API v2. It includes functions to
discover and manage Hue devices, with built-in local bridge interactions.

## Prerequisite

- Get your bridge IP: https://developers.meethue.com/develop/application-design-guidance/hue-bridge-discovery/
- Get your
  Username: [https://developers.meethue.com/develop/hue-api-v2/cloud2cloud-getting-started/](https://developers.meethue.com/develop/hue-api-v2/getting-started/)

## Features

- Interact with Philips Hue API v2
- Discover and manage Hue devices
- Support for both cloud and local bridge interactions
- Cross-platform builds for Linux, Windows, and macOS

## Installation

To install the package, use:

```sh
go get github.com/niomwungeri/hue-v2-api
```

## Download dependencies

```sh
go mod download
```

## Build CLI

```sh
go build -o {build-name}
```

### Sample Client Usage

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

Success sample response:

```json
{
  "errors": [
  ],
  "data": [
    {
      "id": "xxxxx-xxxx-xxxx-acd3-09d8541f903f",
      "product_data": {
        "model_id": "BSB002",
        "manufacturer_name": "Signify Netherlands B.V.",
        "product_name": "Hue Bridge",
        "product_archetype": "bridge_v2",
        "certified": true,
        "software_version": "1.65.1965053040"
      },
      "metadata": {
        "name": "Hue Bridge",
        "archetype": "bridge_v2"
      },
      "identify": {
      },
      "services": [
        {
          "rid": "xxxxx-xxxx-xxxx-acd3-09d8511f903f",
          "rtype": "bridge"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-b091-2b514069d16b",
          "rtype": "zigbee_connectivity"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-8bb3-c1a04aa6397a",
          "rtype": "entertainment"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-a48d-6a4080356c1b",
          "rtype": "zigbee_device_discovery"
        }
      ],
      "type": "device"
    },
    {
      "id": "xxxxx-xxxx-xxxx-0b6d1a91a341",
      "id_v1": "/lights/6",
      "product_data": {
        "model_id": "LOM010",
        "manufacturer_name": "Signify Netherlands B.V.",
        "product_name": "Hue smart plug",
        "product_archetype": "plug",
        "certified": true,
        "software_version": "1.116.3",
        "hardware_platform_type": "100b-11a"
      },
      "metadata": {
        "name": "Device Name",
        "archetype": "plug"
      },
      "identify": {
      },
      "services": [
        {
          "rid": "xxxxx-xxxx-xxxx-aada-8b315679dc47",
          "rtype": "zigbee_connectivity"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-a7eb-32be10cba0f4",
          "rtype": "light"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-a6c5-9bf135b2be8a",
          "rtype": "device_software_update"
        }
      ],
      "type": "device"
    },
    {
      "id": "xxxxx-xxxx-xxxx-82e3-3cf41a6584b0",
      "id_v1": "/lights/4",
      "product_data": {
        "model_id": "LCA009",
        "manufacturer_name": "Signify Netherlands B.V.",
        "product_name": "Hue color lamp",
        "product_archetype": "sultan_bulb",
        "certified": true,
        "software_version": "1.116.3",
        "hardware_platform_type": "100b-114"
      },
      "metadata": {
        "name": "Hue Living lamp ",
        "archetype": "sultan_bulb"
      },
      "identify": {
      },
      "services": [
        {
          "rid": "xxxxx-xxxx-xxxx-a7d7-c591a5639524",
          "rtype": "zigbee_connectivity"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-9c70-848d40336ecc",
          "rtype": "light"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-913b-a82a4e965587",
          "rtype": "entertainment"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-8140-7d7e9ba322c3",
          "rtype": "taurus_7455"
        },
        {
          "rid": "xxxxx-xxxx-xxxx-8140-7d7e9ba322c1",
          "rtype": "device_software_update"
        }
      ],
      "type": "device"
    }
  ]
}
```

# Managing Devices with Hue V2 API

This repository contains a Go client for managing devices through the Hue API v2. The client provides functionalities to get, update, and delete devices. Below are the details and usage instructions for each command.

## Usage

Each client has its own command line for managing devices. Below are the available commands:

### Get Devices

Retrieve a list of devices. To execute this command, use the following syntax:

```sh
./hue-v2-api devices --get --hue-application-key=<your-hue-application-key> --base-url=<your-base-url>
```

Example:

```sh
./hue-v2-api devices --id=6f563b60-1ccd-43a4-83fd-0b6d1a91a341 --get --hue-application-key=<username|hue-application-key> --base-url=https://10.0.0.250
```

Response:
```json
{
  "errors": [],
  "data": [
    {
      "id": "6f563b60-1ccd-43a4-83fd-0b6d1a91a341",
      "id_v1": "/lights/6",
      "product_data": {
        "model_id": "LOM010",
        "manufacturer_name": "Signify Netherlands B.V.",
        "product_name": "Hue smart plug",
        "product_archetype": "plug",
        "certified": true,
        "software_version": "1.116.3",
        "hardware_platform_type": "100b-11a"
      },
      "metadata": {
        "name": "Smart Work Desk",
        "archetype": "plug"
      },
      "identify": {},
      "services": [
        {
          "rid": "3b9bfeb9-41df-4336-aada-8b315679dc47",
          "rtype": "zigbee_connectivity"
        },
        {
          "rid": "8896cd7d-53b9-4037-a7eb-32be10cba0f4",
          "rtype": "light"
        },
        {
          "rid": "e1445ce4-24cb-4f45-a6c5-9bf135b2be8a",
          "rtype": "device_software_update"
        }
      ],
      "type": "device"
    }
  ]
}
```

### Get Device

Retrieve details of a specific device. Use the following syntax:

```sh
./hue-v2-api devices --id=<device-id> --get --hue-application-key=<your-hue-application-key> --base-url=<your-base-url>
```

Example:

```sh
./hue-v2-api devices --id=6f563b60-1ccd-43a4-83fd-0b6d1a91a341 --get --hue-application-key=<username|hue-application-key> --base-url=https://10.0.0.250
```

Response:
```json
{
  "errors": [],
  "data": [
    {
      "id": "6f563b60-1ccd-43a4-83fd-0b6d1a91a341",
      "id_v1": "/lights/6",
      "product_data": {
        "model_id": "LOM010",
        "manufacturer_name": "Signify Netherlands B.V.",
        "product_name": "Hue smart plug",
        "product_archetype": "plug",
        "certified": true,
        "software_version": "1.116.3",
        "hardware_platform_type": "100b-11a"
      },
      "metadata": {
        "name": "Smart Work Desk",
        "archetype": "plug"
      },
      "identify": {},
      "services": [
        {
          "rid": "3b9bfeb9-41df-4336-aada-8b315679dc47",
          "rtype": "zigbee_connectivity"
        },
        {
          "rid": "8896cd7d-53b9-4037-a7eb-32be10cba0f4",
          "rtype": "light"
        },
        {
          "rid": "e1445ce4-24cb-4f45-a6c5-9bf135b2be8a",
          "rtype": "device_software_update"
        }
      ],
      "type": "device"
    }
  ]
}
```

### Update Device

Update the details of a specific device. Use the following syntax:

```sh
./hue-v2-api devices --id=<device-id> --put --hue-application-key=<your-hue-application-key> --base-url=<your-base-url> --payload='<json-payload>'
```

Example:

```sh
./hue-v2-api devices --id=6f563b60-1ccd-43a4-83fd-0b6d1a91a341 --put --hue-application-key=<username|hue-application-key> --base-url=https://10.0.0.250 --payload='{"metadata":{"name":"Smart Work Desk"}}'
```

Response:
```sh
Successfully updated device
```

### Delete Device

Delete a specific device. Use the following syntax:

```sh
./hue-v2-api devices --del --id=<device-id> --hue-application-key=<your-hue-application-key> --base-url=<your-base-url>
```

Example:

```sh
./hue-v2-api devices --del --id=a97e30cf-cfaf-45ff-9e82-58f0c9fd8d97 --hue-application-key=<username|hue-application-key> --base-url=https://10.0.0.250
```

Response:
```json
{
  "errors": [],
  "data": [
    {
      "rid": "a97e30cf-cfaf-45ff-9e82-58f0c9fd8d97",
      "rtype": "device"
    }
  ]
}
```

Feel free to use these commands to manage your Hue devices efficiently. For any issues or further inquiries, please refer to the [documentation](https://developers.meethue.com/develop/hue-api-v2/).

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

Turn the package ready to be used in production, without having to consume the hue endpoints yourself, 
especially if there's not enough time to implement it yourself.

