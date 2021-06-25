package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controller(t *testing.T) {
	c, err := NewController("../display/testdata/acpi_video0")
	assert.Nil(t, err)
	if err != nil {
		return
	}
	brightness, err := c.GetBrightness()
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, brightness, 1)
	list, err := ListControllerPaths("../display/testdata")
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.Equal(t, list, []string{"../display/testdata/acpi_video0", "../display/testdata/intel_backlight"})
}
