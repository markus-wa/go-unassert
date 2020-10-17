# unassert - providing a lack of  assertions

unassert is a library that provides assertions based on the provided [build tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags) - `unassert_panic` or `unassert_stderr`.

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
}
```

### Assertions disabled

```
$ go run main.go
Hello
World
```

### `unassert_panic`

```
$ go run -tags unassert_panic main.go
Hello
panic: UNASSERT: assertion failed: a != b

goroutine 1 [running]:
github.com/markus-wa/go-unassert.Same(...)
    C:/Users/mwalt/dev/go-unassert/panic.go:11
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:14 +0x1c3
exit status 2
```

### `unassert_stderr`

```
$ go run -tags unassert_stderr example/example.go
Hello
UNASSERT: assertion failed: a != b
goroutine 1 [running]:
runtime/debug.Stack(0x4eb280, 0xc000006020, 0xc00009de90)
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:24 +0xa4
github.com/markus-wa/go-unassert/stderr.Stderr(0x4d4798, 0x1a, 0xc0000044a0, 0x2, 0x2)
    C:/Users/mwalt/dev/go-unassert/stderr/stderr.go:19 +0x1ba
github.com/markus-wa/go-unassert.Same(...)
    C:/Users/mwalt/dev/go-unassert/unassert.go:30
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:14 +0x27a
World
UNASSERT: you can also format a string
goroutine 1 [running]:
runtime/debug.Stack(0x4eb280, 0xc000006020, 0xc00009de90)
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:24 +0xa4
github.com/markus-wa/go-unassert/stderr.Stderr(0x4d40fc, 0x18, 0xc000058200, 0x1, 0x1)
    C:/Users/mwalt/dev/go-unassert/stderr/stderr.go:19 +0x1ba
github.com/markus-wa/go-unassert.Error(...)
    C:/Users/mwalt/dev/go-unassert/unassert.go:18
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:18 +0x16d
UNASSERT: you can also pass error messages to any other function
goroutine 1 [running]:
runtime/debug.Stack(0x4eb280, 0xc000006020, 0xc00009de90)
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:24 +0xa4
github.com/markus-wa/go-unassert/stderr.Stderr(0x4d84ee, 0x36, 0x0, 0x0, 0x0)
    C:/Users/mwalt/dev/go-unassert/stderr/stderr.go:19 +0x1ba
github.com/markus-wa/go-unassert.Nilf(...)
    C:/Users/mwalt/dev/go-unassert/unassert.go:98
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:19 +0x19f
UNASSERT: with or without formatting
goroutine 1 [running]:
runtime/debug.Stack(0x4eb280, 0xc000006020, 0xc00009de90)
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:24 +0xa4
github.com/markus-wa/go-unassert/stderr.Stderr(0x4d24a6, 0x12, 0xc000058230, 0x1, 0x1)
    C:/Users/mwalt/dev/go-unassert/stderr/stderr.go:19 +0x1ba
github.com/markus-wa/go-unassert.NotNilf(...)
    C:/Users/mwalt/dev/go-unassert/unassert.go:125
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:20 +0x200
```


## License

This project is licensed under the [MIT license](LICENSE.md).

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmarkus-wa%2Fgo-unassert?ref=badge_large)
