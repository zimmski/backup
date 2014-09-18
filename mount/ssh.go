package mount

import (
	"bytes"
	"fmt"
	"io"

	"github.com/zimmski/backup"
	"github.com/zimmski/backup/exec"
)

type mountSSH struct {
	ssh    string
	folder string

	mounted bool

	password string
	port     uint
}

var _ Mounter = (*mountSSH)(nil)

// NewSSH returns a new SSH mount point
func NewSSH(ssh string, folder string) *mountSSH {
	return &mountSSH{
		ssh:    ssh,
		folder: folder,

		mounted: false,

		password: "",
		port:     0,
	}
}

// attributes

// Password sets the passwort of the connection over STDIN
func (m *mountSSH) Password(set string) {
	m.password = set
}

// Port sets the port of the connection
func (m *mountSSH) Port(set uint) {
	m.port = set
}

// Mounter interface

// Mount mounts this object
func (m *mountSSH) Mount() error {
	if m.mounted {
		return backup.NewError(backup.AlreadyMounted, "mount point already mounted")
	}
	if !backup.FolderExists(m.folder) {
		return backup.NewError(backup.FolderDoesNotExist, fmt.Sprintf("folder %q does not exist", m.folder))
	}

	cmdName := "sshfs"
	var args []string

	args = append(args, m.ssh)
	args = append(args, m.folder)

	if m.port != 0 {
		args = append(args, "-p", fmt.Sprintf("%d", m.port))
	}

	if m.password != "" {
		args = append(args, "-o", "password_stdin")
	}

	cmd := exec.Command(cmdName, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err = cmd.Start()
	if err != nil {
		return err
	}

	if m.password != "" {
		_, err = io.WriteString(stdin, m.password+"\n")
		if err != nil {
			return err
		}

		err = stdin.Close()
		if err != nil {
			return err
		}
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if buf.Len() != 0 {
		return backup.NewError(backup.UnexpectedOutput, fmt.Sprintf("unexpected output: %v", buf.String()))
	}

	m.mounted = true

	return nil
}

// Umount unmounts this object
func (m *mountSSH) Umount() error {
	if !m.mounted {
		return backup.NewError(backup.NotMounted, "mount point needs to be mounted")
	}

	out, err := exec.Combined("umount", m.folder)
	if err != nil {
		return err
	}

	if out != "" {
		return backup.NewError(backup.UnexpectedOutput, fmt.Sprintf("unexpected output: %q", out))
	}

	m.mounted = false

	return nil
}

// AllowAccessToOthers sets the attribute that determines if other users are allowed to access the mount point not just the user who mounted it
func (m *mountSSH) AllowAccessToOthers(set bool) {
	// do nothing
}
