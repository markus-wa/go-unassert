// +build unassert_panic

package unassert

import (
	"github.com/markus-wa/go-unassert/dopanic"
)

const enabled = true

var errorHandler = dopanic.Panic
