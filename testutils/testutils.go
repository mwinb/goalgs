package testutils

import (
	"testing"
)

// ExpectedString Prints test expect vs recieved string
func ExpectedString(t *testing.T, expected string, received string) {
	t.Errorf("Expected %s but got %s", expected, received)
}

// Expect takes in a boolean, and ExpectedString params and executes expected string if true
func Expect(t *testing.T, expected string, received string, passed bool) {
	if !passed {
		ExpectedString(t, expected, received)
	}
}
