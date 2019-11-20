package stderr

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expected, stderr.(*bytes.Buffer).String()[0:len(expected)])

	stderr = old
}
