package main

import (
	"fmt"
	"github.com/niomwungeri-fabrice/hue-v2-api/hue"
)

func main() {
	client, err := hue.NewClient("https://api.meethue.com", "your_bearer_token", "your_hue_application_key")
	if err != nil {
		fmt.Println("error:", err)
	}
	devices, err := client.GetDevices()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, device := range devices {
		fmt.Printf("Device ID: %s, Name: %s, Manufacturer: %s, Model: %s\n", device.Id, device.Metadata.Name, device.ProductData.Manufacturer, device.ProductData.ModelId)
	}
}
