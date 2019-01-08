/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package extfs

import (
	"io"
	"os"
	"time"
)

// File represents a file in the filesystem
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
}

// Filesystem is the filesystem interface.
// Any simulated or real filesystem should implement this interface.
type Filesystem interface {
	Basic
	Dir
	Change
	Closer
}

// Basic abstract the basic operations in a storage-agnostic interface.
type Basic interface {
	// Create creates a file in the filesystem
	Create(filename string) (File, error)

	// Open opens the named file for reading. The associated file descriptor has
	// mode O_RDONLY.
	Open(filename string) (File, error)

	// OpenFile opens a file using the given flags and the given mode.
	OpenFile(filename string, flag int, perm os.FileMode) (File, error)

	// Remove removes the named file.
	Remove(filename string) error

	// RemoveAll removes a directory path and any children it contains.
	// It does not fail if the path does not exist (return nil).
	RemoveAll(path string) error

	// Rename renames (moves) oldpath to newpath. If newpath already exists and
	// is not a directory, Rename replaces it.
	Rename(oldpath, newpath string) error

	// Stat returns a FileInfo describing the named file.
	Stat(filename string) (os.FileInfo, error)
}

// Dir abstract the dir related operations in a storage-agnostic interface.
type Dir interface {
	// ReadDir reads the directory named by dirname and returns a list of
	// directory entries sorted by filename.
	ReadDir(path string) ([]os.FileInfo, error)

	// MkdirAll creates a directory path and all parents that does not exist
	// yet. and returns nil, or else returns an error. If path is
	// already a directory, MkdirAll does nothing and returns nil.
	MkdirAll(path string, perm os.FileMode) error
}

// Change abstract the FileInfo change related operations in a storage-agnostic
// interface.
type Change interface {
	// Chmod changes the mode of the named file to mode.
	Chmod(name string, mode os.FileMode) error

	// Chtimes changes the access and modification times of the named file.
	Chtimes(name string, atime time.Time, mtime time.Time) error
}

// Closer is the interface that wraps the basic Close method.
type Closer interface {
	Close() error
}
