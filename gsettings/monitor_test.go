package gsettings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StartMonitor(t *testing.T) {
	require.Nil(t, StartMonitor())
}
