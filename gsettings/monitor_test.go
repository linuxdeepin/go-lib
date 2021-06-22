package gsettings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StartMonitor(t *testing.T) {
	assert.Nil(t, StartMonitor())
}
