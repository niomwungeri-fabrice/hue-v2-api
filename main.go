package main

import (
	"flag"
	"fmt"
	"github.com/niomwungeri-fabrice/hue-v2-api/cmd"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  hue-v2-api [command] [flags]")
		fmt.Println("")
		fmt.Println("Available Commands:")
		fmt.Println("  --devices     Get list of devices")
		fmt.Println("  --lights      Get list of lights")
		// Add more commands here
		fmt.Println("")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	devicesCommand := flag.NewFlagSet("devices", flag.ExitOnError)
	baseURL := devicesCommand.String("base-url", "https://api.meethue.com", "Base URL for the Hue API")
	bearerToken := devicesCommand.String("bearer-token", "", "Bearer token for the Hue API")
	hueApplicationKey := devicesCommand.String("hue-application-key", "", "Hue application key")

	//lightsCommand := flag.NewFlagSet("lights", flag.ExitOnError)
	//lBaseURL := lightsCommand.String("base-url", "https://api.meethue.com", "Base URL for the Hue API")
	//lBearerToken := lightsCommand.String("bearer-token", "", "Bearer token for the Hue API")
	//lHueApplicationKey := lightsCommand.String("hue-application-key", "", "Hue application key")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--devices":
		devicesCommand.Parse(os.Args[2:])
		cmd.GetDevicesCmd(*baseURL, *bearerToken, *hueApplicationKey)
	case "--lights":
		//lightsCommand.Parse(os.Args[2:])
		//cmd.GetLightsCmd(*lBaseURL, *lBearerToken, *lHueApplicationKey)
	// Add more cases here for additional commands
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}
