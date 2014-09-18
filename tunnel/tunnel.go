package tunnel

// Tunneler defines a generic tunnel
type Tunneler interface {
	// Open opens the tunnel of this object
	Open() error
	// Close closes the tunnel of this object
	Close() error
}
