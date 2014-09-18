// +build example-main

package main

import (
	"github.com/zimmski/backup/backup"
	"github.com/zimmski/backup/mount"
)

func main() {
	ftp := mount.NewFTP("user@localhost/backups", "/mnt/backups")

	ftp.SSL(true)

	err := ftp.Mount()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := ftp.Umount()
		if err != nil {
			panic(err)
		}
	}()

	sync := backup.NewRsync("/important/stuff", "/mnt/backups")

	err = sync.Backup()
	if err != nil {
		panic(err)
	}
}
