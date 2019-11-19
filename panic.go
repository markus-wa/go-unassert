// +build unassert_panic

package unassert

import (
	"fmt"
)

func Same(a interface{}, b interface{}) {
	if a != b {
		panic(fmt.Sprintf("assertion failed: %v != %v", a, b))
	}
}

func True(b bool) {
	if !b {
		panic("assertion failed: value is not true")
	}
}
