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

	errorHandler(format, v...)
}

// True checks if b is true.
// Behaves according to 'unassert_' build tags.
func True(b bool) {
	if !enabled {
		// NOP
		return
	}

	Truef(b, "assertion failed: expected true, got false")
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
		Error(format, v...)
	}
}

// False checks if b is false.
// Behaves according to 'unassert_' build tags.
func False(b bool) {
	if !enabled {
		// NOP
		return
	}

	Falsef(b, "assertion failed: expected false, got true")
}

// Falsef checks if b is false.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Falsef(b bool, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if b {
		Error(format, v...)
	}
}

// Nil checks if v is nil.
// Behaves according to 'unassert_' build tags.
func Nil(o interface{}) {
	if !enabled {
		// NOP
		return
	}

	Nilf(o, "assertion failed: expected nil, got %v", o)
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
		Error(format, v...)
	}
}

// NotNil checks if v is not nil.
// Behaves according to 'unassert_' build tags.
func NotNil(o interface{}) {
	if !enabled {
		// NOP
		return
	}

	NotNilf(o, "assertion failed: expected non-nil, got nil")
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
		Error(format, v...)
	}
}

// Same checks if a == b.
// Behaves according to 'unassert_' build tags.
func Same(a interface{}, b interface{}) {
	if !enabled {
		// NOP
		return
	}

	Samef(a, b, "assertion failed: expected same values, got %v != %v", a, b)
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
		Error(format, v...)
	}
}

// NotSame checks if a != b.
// Behaves according to 'unassert_' build tags.
func NotSame(a interface{}, b interface{}) {
	if !enabled {
		// NOP
		return
	}

	NotSamef(a, b, "assertion failed: expected two different values, got %v", a, b)
}

// NotSamef checks if a != b.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func NotSamef(a interface{}, b interface{}, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if a == b {
		Error(format, v...)
	}
}

// Predicate is a function type used for lazy evaluation in assertions.
// It takes one argument and returns true if that argument matches a given predicate.
// See Matches().
type Predicate func(interface{}) bool

// Matches checks if predicate(v) is true.
// Useful for lazy evaluation of complex assertions.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Matches(predicate Predicate, x interface{}) {
	if !enabled {
		// NOP
		return
	}

	Matchesf(predicate, x, "assertion failed: predicate(x) did not return true; x = %v", x)
}

// Matchesf checks if predicate(v) is true.
// Useful for lazy evaluation of complex assertions.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func Matchesf(predicate Predicate, x interface{}, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if !predicate(x) {
		Error(format, v...)
	}
}

// Evaluator is a function type used for lazy evaluation in assertions.
// Mostly used for function closures.
// See ReturnsTrue().
type Evaluator func() bool

// ReturnsTrue checks if evaluator() returns true.
// Useful for lazy evaluation of complex assertions.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func ReturnsTrue(evaluator Evaluator) {
	if !enabled {
		// NOP
		return
	}

	ReturnsTruef(evaluator, "assertion failed: evaluator() did not return true")
}

// ReturnsTrue checks if evaluator() returns true.
// Useful for lazy evaluation of complex assertions.
// Behaves according to 'unassert_' build tags.
// Formats according to a format specifier.
func ReturnsTruef(evaluator Evaluator, format string, v ...interface{}) {
	if !enabled {
		// NOP
		return
	}

	if !evaluator() {
		Error(format, v...)
	}
}
