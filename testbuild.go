// +build unassert_test

package unassert

const enabled = true

var errorHandler = capture

type errorArgs struct {
	format string
	v      []interface{}
}

var lastError errorArgs

func capture(format string, v ...interface{}) {
	lastError.format = format
	lastError.v = v
}
