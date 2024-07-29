package cmd

import (
	"bytes"
	"fmt"
	"github.com/niomwungeri-fabrice/hue-v2-api/hue"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var (
	getDevice   bool
	delDevice   bool
	putDevice   bool
	deviceID    string
	baseURL     string
	bearerToken string
	hueAppKey   string
	payload     string
)

func getDevicesCmd(baseURL, bearerToken, hueApplicationKey string) {
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
func getDeviceCmd(baseURL, bearerToken, hueApplicationKey, deviceID string) {
	// Create a new Hue API client
	client, err := hue.NewClient(baseURL, bearerToken, hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	devices, err := client.GetDevice(deviceID)
	if err != nil {
		log.Println("error getting device " + deviceID + ": " + err.Error())
		os.Exit(1)
	}

	fmt.Println(hue.JsonConverter(devices))
}
func deleteDevicesCmd(baseURL, bearerToken, hueApplicationKey, deviceID string) {
	// Create a new Hue API client
	client, err := hue.NewClient(baseURL, bearerToken, hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	devices, err := client.DeleteDevice(deviceID)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(hue.JsonConverter(devices))
}
func updateDevicesCmd(baseURL, bearerToken, hueApplicationKey, deviceID string, actions io.Reader) {
	// Create a new Hue API client
	client, err := hue.NewClient(baseURL, bearerToken, hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	err = client.UpdateDevice(deviceID, actions)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

var DevicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Manage devices",
	Long:  "Commands to manage devices",
	Run: func(cmd *cobra.Command, args []string) {
		if getDevice {
			if deviceID == "" {
				getDevicesCmd(baseURL, bearerToken, hueAppKey)
			} else {
				getDeviceCmd(baseURL, bearerToken, hueAppKey, deviceID)
			}
		} else if delDevice {
			if deviceID == "" {
				fmt.Println("error: --id is required for delete")
				os.Exit(1)
			}
			deleteDevicesCmd(baseURL, bearerToken, hueAppKey, deviceID)
		} else if putDevice {
			if deviceID == "" {
				fmt.Println("error: --id is required for update")
				os.Exit(1)
			}
			updateDevicesCmd(baseURL, bearerToken, hueAppKey, deviceID, bytes.NewReader([]byte(payload)))
		} else {
			fmt.Println("error: --get, --del, or --put is required")
			cmd.Usage()
			os.Exit(1)
		}
	},
}

func init() {
	DevicesCmd.Flags().BoolVar(&getDevice, "get", false, "Get devices")
	DevicesCmd.Flags().BoolVar(&delDevice, "del", false, "Delete device")
	DevicesCmd.Flags().BoolVar(&putDevice, "put", false, "Update device")
	DevicesCmd.Flags().StringVar(&deviceID, "id", "", "Device ID")
	DevicesCmd.Flags().StringVar(&baseURL, "base-url", "https://api.meethue.com", "Base URL for the Hue API")
	DevicesCmd.Flags().StringVar(&bearerToken, "bearer-token", "", "Bearer token for the Hue API")
	DevicesCmd.Flags().StringVar(&hueAppKey, "hue-application-key", "", "Hue application key")
	DevicesCmd.Flags().StringVar(&payload, "payload", "", "JSON payload for update a device")

}
