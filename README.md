# backup

This repository holds a set of packages to mount folders, create tunnels and do backups via Go instead of BASH and Perl. External programs are still needed but instead of unmaintainable scripts with confusing arguments and options, clean and structured Go code can be used.

```go
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
```
