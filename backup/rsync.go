package backup

import (
	"github.com/zimmski/backup/exec"
)

type backupRsync struct {
	from string
	to   string

	deleteExtraneous bool
	keepGroup        bool
	keepOwner        bool
	keepPermissions  bool
	tmpFolder        string
}

var _ Backuper = (*backupRsync)(nil)

// NewRsync returns a new rsync backup
func NewRsync(from string, to string) *backupRsync {
	return &backupRsync{
		from: from,
		to:   to,

		deleteExtraneous: false,
		keepGroup:        true,
		keepOwner:        true,
		keepPermissions:  true,
		tmpFolder:        "",
	}
}

// attributes

// DeleteExtraneous sets if extraneous files should be removed from the destination
func (b *backupRsync) DeleteExtraneous(set bool) {
	b.deleteExtraneous = set
}

// KeepGroup sets explicitly that groups should be preserved
func (b *backupRsync) KeepGroup(set bool) {
	b.keepGroup = set
}

// KeepOwner sets explicitly that owners should be preserved
func (b *backupRsync) KeepOwner(set bool) {
	b.keepOwner = set
}

// KeepPermissions sets explicitly that permissions should be preserved
func (b *backupRsync) KeepPermissions(set bool) {
	b.keepPermissions = set
}

// TmpFolder sets the temporary IO folder
func (b *backupRsync) TmpFolder(set string) {
	b.tmpFolder = set
}

// Backuper interface

// Backup starts the backup of this object
func (b *backupRsync) Backup() error {
	cmdName := "rsync"
	var args []string

	args = append(args, "-uav")

	if b.deleteExtraneous {
		args = append(args, "--delete")
	}

	if !b.keepGroup {
		args = append(args, "--no-g")
	}
	if !b.keepOwner {
		args = append(args, "--no-o")
	}
	if !b.keepPermissions {
		args = append(args, "--no-p")
	}

	if b.tmpFolder != "" {
		args = append(args, "--temp-dir="+b.tmpFolder)
	}

	args = append(args, b.from)
	args = append(args, b.to)

	_, err := exec.CombinedWithDirectOutput(cmdName, args...)

	if err != nil {
		return err
	}

	return nil
}
