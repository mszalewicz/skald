package assert

import (
	"fmt"
	"os"
	"runtime/debug"
)

func AssertNil(object any) {
	if object == nil {
		fmt.Fprintf(os.Stderr, "ASSERT\n")
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		os.Exit(1)
	}
}
