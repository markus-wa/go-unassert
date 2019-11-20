// Package unassert provides assertions based on the provided build tags.
// Using -tags unassert_panics panics if an assertion fails.
// Using -tags unassert_stderr prints an error to stdout if an assertion fails.
package unassert

// ErrorHandler handles an error.
// See unassert_panic & unassert_stderr
type ErrorHandler func(format string, v ...interface{})

// Error promotes an error according to 'unassert_' build tags.
// Formats according to a format specifier.
func Error(format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	errorHandler(format, v)
}

// Same checks if a == b.
// Behaves according to 'unassert_' build tags.
func Same(a interface{}, b interface{}) {
	if !enabled {
		// NOP
		return
	}

	if a != b {
		errorHandler("assertion failed: %v != %v", a, b)
	}
}

// Samef checks if a == b.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Samef(a interface{}, b interface{}, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if a != b {
		errorHandler(format, v)
	}
}

// True checks if b is true.
// Behaves according to 'unassert_' build tags.
func True(b bool) {
	if !enabled {
		// NOP
		return
	}

	if !b {
		errorHandler("assertion failed: value is not true")
	}
}

// Truef checks if b is true.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Truef(b bool, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if !b {
		errorHandler(format, v)
	}
}

// Nil checks if v is nil.
// Behaves according to 'unassert_' build tags.
func Nil(o interface{}) {
	if !enabled {
		// NOP
		return
	}

	if o != nil {
		errorHandler("assertion failed: value is not nil")
	}
}

// Nilf checks if v is nil.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Nilf(o interface{}, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if o != nil {
		errorHandler(format, v)
	}
}

// NotNil checks if v is not nil.
// Behaves according to 'unassert_' build tags.
func NotNil(o interface{}) {
	if !enabled {
		// NOP
		return
	}

	if o == nil {
		errorHandler("assertion failed: value is nil")
	}
}

// NotNilf checks if v is not nil.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func NotNilf(o interface{}, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if o == nil {
		errorHandler(format, v)
	}
}
