package utils

import (
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

type UtilsTest struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&UtilsTest{})
}

func (*UtilsTest) TestIsEnvExists(c *C) {
	testEnvName := "test_is_env_exists"
	testEnvValue := "test_env_value"
	c.Check(false, Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, Equals, IsEnvExists(testEnvName))
}

func (*UtilsTest) TestUnsetEnv(c *C) {
	testEnvName := "test_unset_env"
	testEnvValue := "test_env_value"

	c.Check("", Equals, os.Getenv(testEnvName))
	os.Setenv(testEnvName, "")
	c.Check(os.Getenv(testEnvName), Equals, testEnvValue)

	envCount := len(os.Environ())
	UnsetEnv(testEnvName)
	c.Check(os.Getenv(testEnvName), Equals, "")
	c.Check(len(os.Environ()), Equals, envCount-1)
	c.Check(false, Equals, IsEnvExists(testEnvName))
}
