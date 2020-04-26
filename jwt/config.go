package jwt

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// config defines configuration data for JWT token
type config struct {
	Secret string `validate:"required,min=40"`
	Issuer string `validate:"required"`
}

// InitConfig returns a new JWT configuration.
// It also returns an error when the config is invalid
func InitConfig(v *viper.Viper) (*config, error) {
	c := &config{}
	err := v.UnmarshalKey("jwt", c)
	if err != nil {
		return nil, err
	}

	err = c.validate()
	if err != nil {
		return nil, err
	}
	return c, err

}

func (c *config) validate() error {
	v := validator.New()
	return v.Struct(c)
}
