package config

import (
	"testing"
)

func TestConfigurations(t *testing.T) {

	c := GetConfig()
	serverPort := c.Server.Port

	if serverPort != "8080" {
		t.Errorf("Server Port = %s; want 8080", serverPort)
	}

	mockApiUrl := c.MockAPI.Url
	if mockApiUrl != "http://localhost:5000/tax-calculator/tax-year/" {
		t.Errorf("mockApiUrl = %s; want http://localhost:5000/tax-calculator/tax-year/", mockApiUrl)
	}
}
