// +build example-main

package main

import (
	"github.com/zimmski/backup/mount"
)

func main() {
	ftp := mount.NewFTP("user@localhost/backups", "/mnt/backups")

	ftp.SSL(true)

	err := ftp.Mount()
	if err != nil {
		panic(err)
	}
	defer ftp.Umount()

	// use the FTP mount point to backup something
}
