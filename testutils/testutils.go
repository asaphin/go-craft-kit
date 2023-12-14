package testutils

import "testing"

func FailTestIfNil(t *testing.T, value any, message string) {
	if value == nil {
		t.Helper()
		t.Log(message)
		t.Fail()
	}
}
