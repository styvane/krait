package jwt

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
			`[jwt]
Secret = "12345qwerty"
Issuer = "hutsh"
`,
			"weak",
			false,
		},
		{
			`[jwt]
Secret = "aqwAddfdw039945934-1+="
Issuer = ""
`,
			"invalidIssuer",
			false,
		},
		{
			`[jwt]
Issuer = "hutsh"
`,
			"invalidSecret",
			false,
		},
		{
			`[jwt]
Secret = "aqwAddfdw039945934-1+=qwAddfdw039945934-1+="
Issuer = "hut"
`,
			"ok",
			true,
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
				assert.NotNil(t, c)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, c)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	data := []struct {
		secret string
		issuer string
		name   string
		ok     bool
	}{
		{
			"12345qwerty",
			"hutsh",
			"weak",
			false,
		},
		{
			"aqwAddfdw039945934-1+=",
			"",
			"invalidIssuer",
			false,
		},
		{
			"",
			"hutsh",
			"invalidSecret",
			false,
		},
		{
			"aqwAddfdw039945934-1+=qwAddfdw039945934-1+=",
			"hut",
			"ok",
			true,
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{tt.secret, tt.issuer}
			err := c.validate()
			if tt.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}

}
