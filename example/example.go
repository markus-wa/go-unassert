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
