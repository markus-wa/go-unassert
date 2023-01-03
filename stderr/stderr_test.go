package stderr

import (
	"bytes"
	"testing"
)

func TestStderr(t *testing.T) {
	assertStderr(t, func() {
		Stderr("format %s args", "with")
	}, "UNASSERT: format with args")
}

// assertStderr redirects standard output to the stderr variable for tests.
// Panics if an error occurs.
func assertStderr(t *testing.T, f func(), expected string) {
	old := stderr
	stderr = new(bytes.Buffer)

	f()

	if expected != stderr.(*bytes.Buffer).String()[0:len(expected)] {
		t.Errorf("expected %q, got %q", expected, stderr.(*bytes.Buffer).String())
	}

	stderr = old
}
