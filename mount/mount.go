package mount

// Mounter defines a generic mount point
type Mounter interface {
	// Mount mounts this object
	Mount() error
	// Umount unmounts this object
	Umount() error

	// AllowAccessToOthers sets the attribute that determines if other users are allowed to access the mount point not just the user who mounted it
	AllowAccessToOthers(set bool)
}
