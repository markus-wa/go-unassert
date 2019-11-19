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
