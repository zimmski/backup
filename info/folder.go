package info

import (
	"github.com/zimmski/backup/exec"
)

func Folder(path string) (string, error) {
	return exec.Combined("ls", "-l", "-h", path)
}
