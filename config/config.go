package config

import (
	"github.com/spf13/viper"
)

// Config define application configuration
type Config struct {
	Port uint
	Host string
}

// Load create a configuration from a viper instance
func Load(v *viper.Viper) (*Config, error) {
	c := &Config{}
	err := v.Unmarshal(c)
	return c, err
}
