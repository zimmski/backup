package backup

// Backuper defines a generic backup
type Backuper interface {
	// Backup starts the backup of this object
	Backup() error
}
