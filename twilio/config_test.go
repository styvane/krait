package twilio

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	data := []struct {
		sid  string
		tok  string
		ok   bool
		name string
	}{
		{
			"",
			"",
			false,
			"empty",
		},
		{
			"xxxxxxx",
			"xxxx",
			true,
			"ok",
		},
		{
			"xxxxxxx",
			"",
			false,
			"noToken",
		},
		{
			"",
			"xxxxxxx",
			false,
			"noAccount",
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{tt.sid, tt.tok}
			err := c.validate()
			if tt.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}

}

func TestInitConfig(t *testing.T) {
	data := []struct {
		config string
		ok     bool
		name   string
	}{
		{
			"[twilio]",
			false,
			"empty",
		},
		{
			`[twilio]
accountSid = "xxxxxxx"
authToken = "xxxx"
`,
			true,
			"ok",
		},
		{
			`[twilio]
accountSid = "xxxxxxx"
`,
			false,
			"noToken",
		},
		{
			`[twilio]
authToken = "xxxxxxx"
`,
			false,
			"noAccount",
		},
	}

	viper.SetConfigType("toml")
	for _, tt := range data {
		r := strings.NewReader(tt.config)
		t.Run(tt.name, func(t *testing.T) {
			err := viper.ReadConfig(r)
			assert.Nil(t, err)
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
