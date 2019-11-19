# unassert - providing a lack of  assertions

unassert is a library that provides assertions based on the provided [build tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags).

By default assertions are disabled (NOP implementation) to provide the best possible performance.
There are multiple implementations available that can be enabled to discover issues in automated test runs or in QA environments.

[![GoDoc](https://godoc.org/github.com/markus-wa/go-unassert?status.svg)](https://godoc.org/github.com/markus-wa/go-unassert)
[![Build Status](https://travis-ci.com/markus-wa/go-unassert.svg?branch=master)](https://travis-ci.com/markus-wa/go-unassert)
[![codecov](https://codecov.io/gh/markus-wa/go-unassert/branch/master/graph/badge.svg)](https://codecov.io/gh/markus-wa/go-unassert)
[![Go Report](https://goreportcard.com/badge/github.com/markus-wa/go-unassert)](https://goreportcard.com/report/github.com/markus-wa/go-unassert)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE.md)

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
panic: assertion failed: a != b

goroutine 1 [running]:
github.com/markus-wa/go-unassert.Same(...)
    C:/Users/mwalt/dev/go-unassert/panic.go:11
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:14 +0x1c3
exit status 2
```

### `unassert_stderr`

```
$ go run -tags unassert_stderr main.go
Hello
assertion failed: a != b
goroutine 1 [running]:
runtime/debug.Stack(0x19, 0x0, 0x0)
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:24 +0xa4
runtime/debug.PrintStack()
    C:/Users/mwalt/scoop/apps/go/current/src/runtime/debug/stack.go:16 +0x29
github.com/markus-wa/go-unassert.Same(0x4aace0, 0x4e8b00, 0x4aace0, 0x4e8b10)
    C:/Users/mwalt/dev/go-unassert/stderr.go:14 +0xc4
main.main()
    C:/Users/mwalt/dev/go-unassert/example/example.go:14 +0xb6
World

```
