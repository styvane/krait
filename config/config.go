package config

import (
	"github.com/spf13/viper"
)

// JWT defines configuration data for JWT token
type JWT struct {
	Secret string
	Issuer string
}

// TWILIO defines twilio configuration data
type Twilio struct {
	SID   string
	Token string
}

// Application configuration informations
type Config struct {
	Port uint
	Host string
	JWT
	Twilio Twilio
}

// Load create a configuration from a viper instance
func Load(v *viper.Viper) (*Config, error) {
	c := &Config{}
	err := v.Unmarshal(c)
	return c, err
}
