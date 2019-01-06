package extfs

import "errors"

// errors
var (
	ErrCrossedBoundary = errors.New("chroot boundary crossed")
)
