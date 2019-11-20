// +build unassert_stderr

package unassert

import (
	"github.com/markus-wa/go-unassert/stderr"
)

const enabled = true

var errorHandler = stderr.Stderr
