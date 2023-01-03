// +build unassert_panic

package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	defer func() {
		err := recover()

		const expected = "UNASSERT: assertion failed: expected same values, got a != b"

		if err != expected {
			t.Errorf("expected %q, got %q", expected, err)
		}
	}()

	main()
}
