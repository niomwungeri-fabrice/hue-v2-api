package hue

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func isValidBaseURL(url string) bool {
	// Check if the URL starts with "https://"
	if !strings.HasPrefix(url, "https://") {
		return false
	}

	// Extract the IP address part
	ipString := strings.TrimPrefix(url, "https://")

	// Validate the IP address
	ip := net.ParseIP(ipString)
	return ip != nil
}

func JsonConverter(devices interface{}) (string, error) {
	b, err := json.Marshal(devices)
	if err != nil {
		return "", err
	}
	formattedJson := fmt.Sprint(string(b))
	return formattedJson, nil
}
