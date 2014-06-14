package utils

import (
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&Manager{})
}

func (u *Manager) TestUnsetEnv(c *C) {
	testEnvName := "test_env"
	testEnvValue := "test_env_value"

	c.Check("", Equals, os.Getenv(testEnvName))
	os.Setenv(testEnvName, "")
	c.Check(os.Getenv(testEnvName), Equals, testEnvValue)

	envCount := len(os.Environ())
	u.UnsetEnv(testEnvName)
	c.Check(os.Getenv(testEnvName), Equals, "")
	c.Check(len(os.Environ()), Equals, envCount-1)
	foundEnv := false
	for _, e := range os.Environ() {
		if e == testEnvName {
			foundEnv = true
			break
		}
	}
	c.Check(false, Equals, foundEnv)
}
