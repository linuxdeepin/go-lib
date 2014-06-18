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
	c.Check(false, Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, Equals, IsEnvExists(testEnvName))
	envCount := len(os.Environ())
	UnsetEnv(testEnvName)
	c.Check(false, Equals, IsEnvExists(testEnvName))
	c.Check(len(os.Environ()), Equals, envCount-1)
}

func (*UtilsTest) TestUnsetEnvC(c *C) {
	testEnvName := "test_unset_envc"
	testEnvValue := "test_env_value"
	c.Check(false, Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, Equals, IsEnvExists(testEnvName))
	envCount := len(os.Environ())
	UnsetEnvC(testEnvName)
	c.Check(false, Equals, IsEnvExists(testEnvName))
	c.Check(len(os.Environ()), Equals, envCount-1)
}

func (*UtilsTest) TestClearEnvC(c *C) {
	ClearEnvC()
	c.Check(len(os.Environ()), Equals, 0)
}
