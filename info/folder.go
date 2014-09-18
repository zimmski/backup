package info

import (
	"github.com/zimmski/backup/exec"
)

// Folder returns the output of a verbose ls of a folder
func Folder(path string) (string, error) {
	return exec.Combined("ls", "-l", "-h", path)
}
