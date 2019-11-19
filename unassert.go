// +build !unassert_stderr,!unassert_panic

package unassert

// Same checks if a == b.
// Behaves according to 'unassert_' build tags.
func Same(a interface{}, b interface{}) {
	// NOP
}

// Same checks if b is true.
// Behaves according to 'unassert_' build tags.
func True(b bool) {
	// NOP
}
