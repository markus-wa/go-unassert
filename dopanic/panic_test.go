package dopanic_test

import (
	"testing"

	"github.com/markus-wa/go-unassert/dopanic"
)

func TestPanic(t *testing.T) {
	defer func() {
		err := recover()

		const expected = "UNASSERT: format with args"

		if err != expected {
			t.Errorf("expected %q, got %q", expected, err)
		}
	}()

	dopanic.Panic("format %s args", "with")
}
