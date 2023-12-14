package testutils

import "testing"

func TestFailTestIfNil(t *testing.T) {
	t.Skip()

	FailTestIfNil(t, nil, "should not be nil")
}
