package output

import (
	"fmt"
	"os"
)

// OnError prints an error if err is not nil and exits the program.
func OnError(err error, text string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, text+": %s", err.Error())
		os.Exit(1)
	}
}

// Error prints an error and exits the program.
func Error(text string) {
	fmt.Fprintln(os.Stderr, text)
	os.Exit(1)
}
