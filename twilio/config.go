package twilio

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// config defines twilio configuration data
type config struct {
	AccountSid string `validate:"required"`
	AuthToken  string `validate:"required"`
}

func (c *config) validate() error {
	v := validator.New()
	return v.Struct(c)
}

func InitConfig(v *viper.Viper) (*config, error) {
	c := &config{}
	err := v.UnmarshalKey("twilio", c)
	if err != nil {
		return nil, err
	}

	err = c.validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}
