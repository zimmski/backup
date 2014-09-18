# backup

This repository holds a set of packages to mount and create backups via Go instead of BASH and Perl. External programs are still needed but instead of unmaintainable scripts, clean and structured Go code can be used.

```go
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
```
