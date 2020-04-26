package mongo

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	data := []struct {
		config string
		name   string
		ok     bool
	}{
		{
			`[mongodb]
uri = "mongodb://localhost:27017"
`,
			"ok",
			true,
		},
		{
			`[mongodb]
uri = "localhost"
`,
			"nok",
			false,
		},
	}

	viper.SetConfigType("toml")
	for _, tt := range data {
		r := strings.NewReader(tt.config)
		err := viper.ReadConfig(r)
		assert.Nil(t, err)
		t.Run(tt.name, func(t *testing.T) {
			c, err := InitConfig(viper.GetViper())
			if tt.ok {
				assert.Nil(t, err)
				assert.NotNil(t, c)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, c)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	data := []struct {
		uri  string
		name string
		ok   bool
	}{
		{
			"mongodb://localhost:27017",
			"ok",
			true,
		},
		{
			"localhost",
			"nok",
			false,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{tt.uri}
			err := c.validate()
			if tt.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}

}
