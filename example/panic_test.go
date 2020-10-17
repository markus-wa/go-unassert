// +build unassert_panic

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	defer func() {
		err := recover()

		assert.Equal(t, "UNASSERT: assertion failed: expected same values, got a != b", err)
	}()

	main()
}
