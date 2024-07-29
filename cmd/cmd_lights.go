package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/niomwungeri-fabrice/hue-v2-api/hue"
)

func GetLightsCmd(baseURL, bearerToken, hueApplicationKey string) {
	// Create a new Hue API client
	client, err := hue.NewClient(baseURL, bearerToken, hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lights, err := client.GetLights()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(hue.JsonConverter(lights))
}

func InitLight() *cobra.Command {
	DevicesCmd.Flags().BoolVar(&getDevice, "get", false, "Get devices")
	DevicesCmd.Flags().BoolVar(&delDevice, "del", false, "Delete device")
	DevicesCmd.Flags().BoolVar(&putDevice, "put", false, "Update device")
	DevicesCmd.Flags().StringVar(&deviceID, "id", "", "Device ID")
	DevicesCmd.Flags().StringVar(&baseURL, "base-url", "https://api.meethue.com", "Base URL for the Hue API")
	DevicesCmd.Flags().StringVar(&bearerToken, "bearer-token", "", "Bearer token for the Hue API")
	DevicesCmd.Flags().StringVar(&hueAppKey, "hue-application-key", "", "Hue application key")
	return DevicesCmd
}
