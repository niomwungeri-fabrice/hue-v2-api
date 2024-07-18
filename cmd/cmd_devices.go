package cmd

import (
	"fmt"
	"os"

	"github.com/niomwungeri-fabrice/hue-v2-api/hue"
)

func GetDevicesCmd(baseURL, bearerToken, hueApplicationKey string) {
	// Create a new Hue API client
	client, err := hue.NewClient(baseURL, bearerToken, hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	devices, err := client.GetDevices(false)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(hue.JsonConverter(devices))
}
