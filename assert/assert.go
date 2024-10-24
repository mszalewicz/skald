package assert

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Assert(condition bool) {
	if !condition {
		fmt.Fprintf(os.Stderr, "ASSERT\n")
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		os.Exit(1)
	}
}

func AssertNil(object any) {
	if object == nil {
		fmt.Fprintf(os.Stderr, "ASSERT\n")
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		os.Exit(1)
	}
}

func Maybe(object any) {

}
