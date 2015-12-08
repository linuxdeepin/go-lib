package iso

import (
	C "launchpad.net/gocheck"
	"testing"
)

type testWrapper struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&testWrapper{})
}
