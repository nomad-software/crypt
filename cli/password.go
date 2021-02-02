package cli

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadPassword reads a password from the terminal
func ReadPassword(prompt string) ([]byte, error) {
	fmt.Fprint(os.Stderr, prompt)
	var fd int
	if terminal.IsTerminal(syscall.Stdin) {
		fd = syscall.Stdin
	} else {
		tty, err := os.Open("/dev/tty")
		if err != nil {
			return nil, err
		}
		defer tty.Close()
		fd = int(tty.Fd())
	}

	pass, err := terminal.ReadPassword(fd)
	fmt.Fprintln(os.Stderr)
	return pass, err
}
