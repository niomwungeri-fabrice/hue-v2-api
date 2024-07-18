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

	commands := flag.NewFlagSet("devices", flag.ExitOnError)
	baseURL := commands.String("base-url", "https://api.meethue.com", "Base URL for the Hue API")
	bearerToken := commands.String("bearer-token", "", "Bearer token for the Hue API")
	hueApplicationKey := commands.String("hue-application-key", "", "Hue application key")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--devices":
		commands.Parse(os.Args[2:])
		cmd.GetDevicesCmd(*baseURL, *bearerToken, *hueApplicationKey)
	case "--lights":
		commands.Parse(os.Args[2:])
		cmd.GetLightsCmd(*baseURL, *bearerToken, *hueApplicationKey)
	// Add more cases here for additional commands
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}
