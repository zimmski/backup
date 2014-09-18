package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	osexec "os/exec"
	"strings"

	"github.com/zimmski/backup"
)

// Combined executes a command with given arguments and returns the combined output
func Combined(name string, args ...string) (string, error) {
	if backup.Verbose {
		fmt.Fprintln(os.Stderr, "Execute: ", name, strings.Join(args, " "))
	}

	cmd := osexec.Command(name, args...)

	out, err := cmd.CombinedOutput()

	return string(out), err
}

// CombinedWithDirectOutput executes a command with given arguments and prints (to StdOut) and returns the combined output
func CombinedWithDirectOutput(name string, args ...string) (string, error) {
	if backup.Verbose {
		fmt.Fprintln(os.Stderr, "Execute: ", name, strings.Join(args, " "))
	}

	cmd := osexec.Command(name, args...)

	var buf bytes.Buffer

	out := io.MultiWriter(os.Stdout, &buf)

	cmd.Stderr = out
	cmd.Stdout = out

	err := cmd.Run()

	return buf.String(), err
}

// Command returns a generic exec command
func Command(name string, args ...string) *osexec.Cmd {
	if backup.Verbose {
		fmt.Fprintln(os.Stderr, "Execute: ", name, strings.Join(args, " "))
	}

	return osexec.Command(name, args...)
}
