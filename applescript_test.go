package applescript_test

import (
	"testing"

	"github.com/josa42/go-applescript"
)

func assertFail(t *testing.T, expected string, fn func() (string, error)) {

	actual, err := fn()

	if err == nil {
		t.Errorf("error should not be nil")
	}

	if expected != "" && actual != expected {
		t.Errorf("output should be '%s'", expected)
	}
}
func assertSuccess(t *testing.T, expected string, fn func() (string, error)) {

	actual, err := fn()

	if err != nil {
		t.Errorf("error should be nil (is: %s)", err.Error())
	}

	if expected != "" && actual != expected {
		t.Errorf("output should be '%s'", expected)
	}
}

func Test_Run(t *testing.T) {
	assertFail(t, "", func() (string, error) {
		return applescript.Run("applcation Safari")
	})
}

func Test_Run_fail(t *testing.T) {
	assertSuccess(t, "Test", func() (string, error) {
		return applescript.Run(`set test to "Test"`)
	})
}

func Test_RunScript(t *testing.T) {
	assertSuccess(t, "Test", func() (string, error) {
		return applescript.RunScript("testdata/test.applescript")
	})
}

func Test_RunScript2(t *testing.T) {
	assertSuccess(t, "Test", func() (string, error) {
		return applescript.RunScript("testdata/test.scpt")
	})
}
