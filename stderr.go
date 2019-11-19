// +build unassert_stderr

package unassert

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

var stderr io.Writer = os.Stderr

func Same(a interface{}, b interface{}) {
	if a != b {
		fmt.Fprintf(stderr, "assertion failed: %v != %v\n", a, b)
		stderr.Write(debug.Stack())
	}
}

func True(b bool) {
	if !b {
		fmt.Fprintln(stderr, "assertion failed: value is not true")
		stderr.Write(debug.Stack())
	}
}
