package mongo

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type config struct {
	URI string `validate:"required,uri"`
}

// InitConfig returns a new valid MongoDB client configuration
// It also returns an error when the configuration is invalid
func InitConfig(v *viper.Viper) (*config, error) {
	c := &config{}
	err := v.UnmarshalKey("mongodb", c)
	if err != nil {
		return nil, err
	}

	err = c.validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *config) validate() error {
	v := validator.New()
	return v.Struct(c)
}
