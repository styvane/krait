package server

import (
	"testing"

	"github.com/hutsharing/krait/config"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := New(&config.Config{})
	assert.NotNil(t, s)

}
