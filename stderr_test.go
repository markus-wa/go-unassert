// +build unassert_stderr

package unassert

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSame(t *testing.T) {
	assertStderr(t, func() {
		Same(1, 1)
	}, "")
}

func TestSame_Stderr(t *testing.T) {
	assertStderr(t, func() {
		Same(1, 2)
	}, "assertion failed: 1 != 2")
}

func TestTrue(t *testing.T) {
	assertStderr(t, func() {
		True(true)
	}, "")
}

func TestTrue_Stderr(t *testing.T) {
	assertStderr(t, func() {
		True(false)
	}, "assertion failed: value is not true")
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
