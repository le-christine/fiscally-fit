package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server  ServerConfigurations
	MockAPI MockAPIConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port string
}

// MockAPIConfigurations exported
type MockAPIConfigurations struct {
	Url string
}

var Configuration Configurations

func init() {
	fmt.Println("Initializing the configurations")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Config initialized")
}

func GetConfig() Configurations {
	fmt.Println(Configuration)
	return Configuration
}
