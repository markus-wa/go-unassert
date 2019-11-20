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
