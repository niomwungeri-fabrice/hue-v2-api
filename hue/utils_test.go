package hue

import "testing"

func TestIsValidBaseURL(t *testing.T) {
	tests := []struct {
		baseURL string
		valid   bool
	}{
		{"https://192.168.1.1", true},
		{"https://127.0.0.1", true},
		{"https://255.255.255.255", true},
		{"https://0.0.0.0", true},
		{"http://192.168.1.1", false}, // Invalid: should start with https://
		{"https://192.168.1", false},  // Invalid: incomplete IP address
		{"192.168.1.1", false},        // Invalid: no https://
	}

	for _, tt := range tests {
		t.Run(tt.baseURL, func(t *testing.T) {
			if got := isValidBaseURL(tt.baseURL); got != tt.valid {
				t.Errorf("isValidBaseURL(%v) = %v, want %v", tt.baseURL, got, tt.valid)
			}
		})
	}
}
