package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/niomwungeri-fabrice/hue-v2-api/hue"
)

func main() {

	// Define command-line flags with default values
	baseURL := flag.String("base-url", "https://api.meethue.com", "Base URL for the Hue API")
	bearerToken := flag.String("bearer-token", "", "Bearer token for the Hue API")
	hueApplicationKey := flag.String("hue-application-key", "", "Hue application key")

	// Parse the flags
	flag.Parse()

	if *hueApplicationKey == "" {
		fmt.Println("Warning: --hue-application-key is not provided, some features may not work correctly")
	}

	// Validate the baseURL and hueApplicationKey
	if *baseURL == "" {
		fmt.Println("Error: --base-url is required")
		flag.Usage()
		os.Exit(1)
	}

	// Create a new Hue API client
	client, err := hue.NewClient(*baseURL, *bearerToken, *hueApplicationKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	devices, err := client.GetDevices(false)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(hue.JsonConverter(devices))
}
