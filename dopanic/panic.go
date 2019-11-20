// Package dopanic implements unassert.ErrorHandler by panicking.
package dopanic

import (
	"fmt"
)

// Panic formats an error and panics.
func Panic(format string, v ...interface{}) {
	panic("UNASSERT: " + fmt.Sprintf(format, v...))
}
