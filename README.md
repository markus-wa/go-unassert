# unassert - providing a lack of  assertions

unassert is a library that provides assertions based on the provided [build tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags) - [`unassert_panic`](#unassert_panic) or [`unassert_stderr`](#unassert_stderr).

By default assertions are disabled (NOP implementation) to provide zero overhead in performance.
The Go compiler detects when the library is not enbled and completely removes any function calls to `unassert`, so it behaves as if it the library were not used at all.
This can be very useful if you are in a high performance scenario where you want to do assertinons in tight loops and you are reasonably sure your test cases cover any realistic scenarios.

There are multiple implementations available that can be enabled to discover issues in automated test runs or in QA environments.

[![GoDoc](https://godoc.org/github.com/markus-wa/go-unassert?status.svg)](https://godoc.org/github.com/markus-wa/go-unassert)
[![Build Status](https://travis-ci.com/markus-wa/go-unassert.svg?branch=master)](https://travis-ci.com/markus-wa/go-unassert)
[![codecov](https://codecov.io/gh/markus-wa/go-unassert/branch/master/graph/badge.svg)](https://codecov.io/gh/markus-wa/go-unassert)
[![Go Report](https://goreportcard.com/badge/github.com/markus-wa/go-unassert)](https://goreportcard.com/report/github.com/markus-wa/go-unassert)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.md)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert?ref=badge_shield)

## Go Get

    go get -u github.com/markus-wa/go-unassert

## Example

```go
package main

import (
	"fmt"

	"github.com/markus-wa/go-unassert"
)

func main() {
	unassert.True(true)

	fmt.Println("Hello")

	unassert.Same("a", "b")

	fmt.Println("World")

	unassert.Error("you can also %s a string", "format")
	unassert.Nilf(new(struct{}), "you can also pass error messages to any other function")
	unassert.NotNilf(nil, "with or without %s", "formatting")

	lenLessThan10 := func(v interface{}) bool {
		return len(v.([]int)) < 10
	}
	unassert.Matches(lenLessThan10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	complexEvaluation := func() bool {
		// add logic here
		return false
	}
	unassert.ReturnsTrue(complexEvaluation)
}
```

### Assertions disabled

    go run main.go

output:

```
Hello
World
```

### `unassert_panic`

    go run -tags unassert_panic main.go

output:

```
Hello
panic: UNASSERT: assertion failed: expected same values, got a != b

goroutine 1 [running]:
github.com/markus-wa/go-unassert.Same(...)
    C:/Users/mwalt/dev/go-unassert/panic.go:11
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:14 +0x1c3
exit status 2
```

### `unassert_stderr`

    go run -tags unassert_stderr example/example.go

output:

```
Hello
UNASSERT: assertion failed: expected same values, got a != b
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084ea8)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c8eef, 0x34, 0xc00000c060, 0x2, 0x2)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
github.com/markus-wa/go-unassert.Samef(...)
    /home/mark/dev/go-unassert/unassert.go:142
github.com/markus-wa/go-unassert.Same(...)
    /home/mark/dev/go-unassert/unassert.go:129
main.main()
    /home/mark/dev/go-unassert/example/example.go:14 +0x3c4
World
UNASSERT: you can also format a string
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084ea8)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c51f2, 0x18, 0xc000010210, 0x1, 0x1)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
main.main()
    /home/mark/dev/go-unassert/example/example.go:18 +0x18f
UNASSERT: you can also pass error messages to any other function
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084ea8)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c9136, 0x36, 0x0, 0x0, 0x0)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
github.com/markus-wa/go-unassert.Nilf(...)
    /home/mark/dev/go-unassert/unassert.go:92
main.main()
    /home/mark/dev/go-unassert/example/example.go:19 +0x1c2
UNASSERT: with or without formatting
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084ea8)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c3db5, 0x12, 0xc000010240, 0x1, 0x1)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
github.com/markus-wa/go-unassert.NotNilf(...)
    /home/mark/dev/go-unassert/unassert.go:117
main.main()
    /home/mark/dev/go-unassert/example/example.go:20 +0x224
UNASSERT: assertion failed: predicate(x) did not return true; x = [1 2 3 4 5 6 7 8 9 10]
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084e70)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c93d2, 0x3a, 0xc000010260, 0x1, 0x1)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
github.com/markus-wa/go-unassert.Matchesf(0x4c9920, 0x49fb40, 0xc00000c080, 0x4c93d2, 0x3a, 0xc000010260, 0x1, 0x1)
    /home/mark/dev/go-unassert/unassert.go:200 +0x88
github.com/markus-wa/go-unassert.Matches(...)
    /home/mark/dev/go-unassert/unassert.go:186
main.main()
    /home/mark/dev/go-unassert/example/example.go:25 +0x328
UNASSERT: assertion failed: evaluator() did not return true
goroutine 1 [running]:
runtime/debug.Stack(0x4ddcc0, 0xc00000e020, 0xc000084e70)
    /opt/go/current/src/runtime/debug/stack.go:24 +0x9d
github.com/markus-wa/go-unassert/stderr.Stderr(0x4c8c09, 0x31, 0x0, 0x0, 0x0)
    /home/mark/dev/go-unassert/stderr/stderr.go:20 +0x1b3
github.com/markus-wa/go-unassert.Error(...)
    /home/mark/dev/go-unassert/unassert.go:18
github.com/markus-wa/go-unassert.ReturnsTruef(0x4c9928, 0x4c8c09, 0x31, 0x0, 0x0, 0x0)
    /home/mark/dev/go-unassert/unassert.go:233 +0x74
github.com/markus-wa/go-unassert.ReturnsTrue(...)
    /home/mark/dev/go-unassert/unassert.go:219
main.main()
    /home/mark/dev/go-unassert/example/example.go:31 +0x35e
```

## License

This project is licensed under the [MIT license](LICENSE.md).

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert?ref=badge_large)
