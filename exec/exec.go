package exec

import (
	osexec "os/exec"
)

// Combined executes a command with given arguments and returns the combined output
func Combined(name string, args ...string) (string, error) {
	cmd := osexec.Command(name, args...)

	out, err := cmd.CombinedOutput()

	return string(out), err
}
