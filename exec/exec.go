package exec

import (
	"bytes"
	"io"
	"os"
	osexec "os/exec"
)

// Combined executes a command with given arguments and returns the combined output
func Combined(name string, args ...string) (string, error) {
	cmd := osexec.Command(name, args...)

	out, err := cmd.CombinedOutput()

	return string(out), err
}

func CombinedWithDirectOutput(name string, args ...string) (string, error) {
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
	return osexec.Command(name, args...)
}
