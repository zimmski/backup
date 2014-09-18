package tunnel

import (
	"fmt"
	"os/exec"

	"github.com/zimmski/backup"
)

type tunnelSSH struct {
	over string
	to   string
	port uint

	opened bool
	cmd    *exec.Cmd

	compress bool
}

var _ Tunneler = (*tunnelSSH)(nil)

// New returns a new SSH tunnel
func NewSSH(over string, to string, onPort uint) *tunnelSSH {
	return &tunnelSSH{
		over: over,
		to:   to,
		port: onPort,

		opened: false,
		cmd:    nil,

		compress: false,
	}
}

// attributes

// Compress sets if the tunnel should be compressed
func (t *tunnelSSH) Compress(set bool) {
	t.compress = set
}

// Tunneler interface

// Open opens the tunnel of this object
func (t *tunnelSSH) Open() error {
	if t.opened {
		return backup.NewError(backup.AlreadyOpened, "tunnel already opened")
	}

	cmdName := "ssh"
	var args []string

	args = append(args, t.over)

	if t.compress {
		// compress connection
		args = append(args, "-C")
	}

	// bind port to address
	args = append(args, "-L", fmt.Sprintf("%d:%s", t.port, t.to))

	// do not execute any commands
	args = append(args, "-N")

	// let ssh go into the background
	args = append(args, "-f")

	t.cmd = exec.Command(cmdName, args...)

	err := t.cmd.Run()
	if err != nil {
		return err
	}

	// TODO check stdout and stderr

	t.opened = true

	return nil
}

// Close closes the tunnel of this object

func (t *tunnelSSH) Close() error {
	if !t.opened {
		return backup.NewError(backup.NotOpened, "tunnel needs to be opened")
	}

	err := t.cmd.Process.Kill()
	if err != nil {
		return err
	}

	t.cmd = nil
	t.opened = false

	return nil
}
