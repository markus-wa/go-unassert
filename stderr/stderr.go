// Package stderr implements unassert.ErrorHandler by printing to stderr.
package stderr

import (
	"fmt"
	"io"
	"os"
)

var stderr io.Writer = os.Stderr

// Stderr prints an error to stderr.
func Stderr(format string, v ...interface{}) {
	_, err := fmt.Fprintf(stderr, format, v...)
	fmt.Println("failed to write to stderr:", err)
}
