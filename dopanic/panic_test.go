package dopanic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/markus-wa/go-unassert/dopanic"
)

func TestPanic(t *testing.T) {
	defer func() {
		err := recover()

		assert.Equal(t, "format with args", err)
	}()

	dopanic.Panic("format %s args", "with")
}
