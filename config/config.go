package config

import "github.com/spf13/viper"

type Config struct {
	JWTSecret []byte `json: "jwt.secret"`
	JWTIssuer string `json: "jwt.issuer"`
}

func Load(v *viper.Viper) *Config {
	c := &Config{}
	v.Unmarshal(c)
	return c
}
