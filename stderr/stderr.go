// Package stderr implements unassert.ErrorHandler by printing to stderr.
package stderr

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

var stderr io.Writer = os.Stderr

// Stderr prints an error to stderr.
func Stderr(format string, v ...interface{}) {
	_, err := fmt.Fprintln(stderr, fmt.Sprintf("UNASSERT: "+format, v...))
	if err != nil {
		fmt.Println("failed to write assertion error to stderr:", err)
	}
	_, err = stderr.Write(debug.Stack())
	if err != nil {
		fmt.Println("failed to write stack to stderr:", err)
	}
}
