package assert

import (
	"fmt"
	"os"
	"runtime/debug"
)

const anotationLen = 110

func annotate(text string) {
	if len(text) > 0 {
		fmt.Fprintf(os.Stderr, "\n____ [ ")
		fmt.Fprint(os.Stderr, text)
		fmt.Fprintf(os.Stderr, " ] ")
		for range anotationLen - 10 - len(text) {
			fmt.Fprintf(os.Stderr, "_")
		}
	} else {
		for range anotationLen {
			fmt.Fprintf(os.Stderr, "_")
		}
	}
	fmt.Fprintf(os.Stderr, "\n\n")
}

func Assert(condition bool) {
	if !condition {
		annotate("ASSERT")
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		annotate("")
		os.Exit(1)
	}
}

func NotNil(object any) {
	if object == nil {
		annotate("ASSERT_NIL")
		fmt.Fprintln(os.Stderr, string(debug.Stack()))
		annotate("")
		os.Exit(1)
	}
}

func Maybe(object any) {

}
