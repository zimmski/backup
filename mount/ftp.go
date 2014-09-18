package mount

import (
	"fmt"

	"github.com/zimmski/backup"
	"github.com/zimmski/backup/exec"
)

type mountFTP struct {
	ftp    string
	folder string

	mounted bool

	allowAccessToOthers bool

	ssl            bool
	sslDoNotVerify bool
}

var _ Mounter = (*mountFTP)(nil)

// NewFTP returns a new FTP mount point
func NewFTP(ftp string, folder string) *mountFTP {
	return &mountFTP{
		ftp:    ftp,
		folder: folder,

		mounted: false,

		ssl: false,
	}
}

// attributes

// SSL sets if the FTP connection uses SSL
func (m *mountFTP) SSL(set bool) {
	m.ssl = set
}

// SSLDoNotVerify sets if SSL connections are verified, like the credibility of the certificate
func (m *mountFTP) SSLDoNotVerify(set bool) {
	m.sslDoNotVerify = set
}

// Mounter interface

// Mount mounts this object
func (m *mountFTP) Mount() error {
	if m.mounted {
		return backup.NewError(backup.AlreadyMounted, "mount point already mounted")
	}
	if !backup.FolderExists(m.folder) {
		return backup.NewError(backup.FolderDoesNotExist, fmt.Sprintf("folder %q does not exist", m.folder))
	}

	cmdName := "curlftpfs"
	var args []string

	if m.allowAccessToOthers {
		args = append(args, "-o", "allow_other")
	}

	if m.ssl {
		args = append(args, "-o", "ssl")

		if m.sslDoNotVerify {
			args = append(args, "-o", "no_verify_peer")
		}
	}

	args = append(args, m.ftp)
	args = append(args, m.folder)

	out, err := exec.Combined(cmdName, args...)
	if err != nil {
		return err
	}

	if out != "" {
		return backup.NewError(backup.UnexpectedOutput, fmt.Sprintf("unexpected output: %q", out))
	}

	m.mounted = true

	return nil
}

// Umount unmounts this object
func (m *mountFTP) Umount() error {
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
func (m *mountFTP) AllowAccessToOthers(set bool) {
	m.allowAccessToOthers = set
}
