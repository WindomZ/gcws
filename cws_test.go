package gcws

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestRegister(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	Register("", nil)
}

func TestNewCWS(t *testing.T) {
	_, err := NewCWS("")
	assert.Error(t, err)
}
