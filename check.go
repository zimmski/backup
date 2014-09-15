package backup

import (
	"os"
)

// FolderExists checks if a folder exists
func FolderExists(folder string) bool {
	stat, err := os.Stat(folder)
	if err != nil {
		return false
	}

	return stat.IsDir()
}
