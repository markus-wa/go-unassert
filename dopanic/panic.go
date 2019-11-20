package dopanic

import (
	"fmt"
)

func Panic(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}
