// +build unassert_panic

package unassert_test

import (
	"testing"

	"github.com/markus-wa/go-unassert"

	"github.com/stretchr/testify/assert"
)

func TestSame(t *testing.T) {
	unassert.Same(1, 1)
}

func TestSame_Panic(t *testing.T) {
	defer func() {
		err := recover()

		assert.Equal(t, "assertion failed: 1 != 2", err)
	}()

	unassert.Same(1, 2)
}

func TestTrue(t *testing.T) {
	unassert.True(true)
}

func TestTrue_Panic(t *testing.T) {
	defer func() {
		err := recover()

		assert.Equal(t, "assertion failed: value is not true", err)
	}()

	unassert.True(false)
}
