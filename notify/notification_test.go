package notify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_notification(t *testing.T) {
	Init("test")
	assert.True(t, IsInitted())
	SetAppName("testNotify")
	assert.Equal(t, GetAppName(), "testNotify")

	notify := NewNotification("", "test", "notification-bluetooth-connected")
	assert.Equal(t, notify.getAppName(), "testNotify")

	notify.Update("", "testNotify", "notification-bluetooth-connected")
	notify.Destroy()
	Destroy()
}
