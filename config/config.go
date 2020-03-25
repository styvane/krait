package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// JWT defines configuration data for JWT token
type JWT struct {
	Secret string
	Issuer string
}

// Application configuration informations
type Config struct {
	JWT
	Port uint
	Host string
}

// Load create a configuration from a viper instance
func Load(v *viper.Viper) (*Config, error) {
	c := &Config{}
	err := v.Unmarshal(c)
	fmt.Println(c)
	return c, err
}
