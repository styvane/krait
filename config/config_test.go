package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	_, err := Load(viper.GetViper())
	assert.Nil(t, err)
}
