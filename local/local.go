/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package local

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rkcloudchain/extfs"
)

const (
	defaultDirectoryMode = 0755
	defaultCreateMode    = 0666
)

// Local is a filesystem based on the local filesystem.
type Local struct {
	base string
}

// New returns a local filesystem.
func New(baseDir string) extfs.Filesystem {
	return &Local{baseDir}
}

// Create ...
func (fs *Local) Create(filename string) (extfs.File, error) {
	fullpath, err := fs.underlyingPath(filename)
	if err != nil {
		return nil, err
	}

	return fs.openFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, defaultCreateMode)
}

// Open ...
func (fs *Local) Open(filename string) (extfs.File, error) {
	fullpath, err := fs.underlyingPath(filename)
	if err != nil {
		return nil, err
	}

	return os.Open(fullpath)
}

// OpenFile ...
func (fs *Local) OpenFile(filename string, flag int, perm os.FileMode) (extfs.File, error) {
	fullpath, err := fs.underlyingPath(filename)
	if err != nil {
		return nil, err
	}

	return fs.openFile(fullpath, flag, perm)
}

// Rename ...
func (fs *Local) Rename(from, to string) error {
	var err error
	from, err = fs.underlyingPath(from)
	if err != nil {
		return err
	}

	to, err = fs.underlyingPath(to)
	if err != nil {
		return err
	}

	if err := fs.createDir(to); err != nil {
		return err
	}

	return os.Rename(from, to)
}

// Remove ...
func (fs *Local) Remove(filename string) error {
	fullpath, err := fs.underlyingPath(filename)
	if err != nil {
		return err
	}

	return os.Remove(fullpath)
}

// RemoveAll ...
func (fs *Local) RemoveAll(path string) error {
	fullpath, err := fs.underlyingPath(path)
	if err != nil {
		return err
	}

	return os.RemoveAll(fullpath)
}

// ReadDir ...
func (fs *Local) ReadDir(path string) ([]os.FileInfo, error) {
	fullpath, err := fs.underlyingPath(path)
	if err != nil {
		return nil, err
	}

	l, err := ioutil.ReadDir(fullpath)
	if err != nil {
		return nil, err
	}

	return l[:], nil
}

// MkdirAll ...
func (fs *Local) MkdirAll(path string, perm os.FileMode) error {
	fullpath, err := fs.underlyingPath(path)
	if err != nil {
		return err
	}

	return os.MkdirAll(fullpath, defaultDirectoryMode)
}

// Stat ...
func (fs *Local) Stat(filename string) (os.FileInfo, error) {
	fullpath, err := fs.underlyingPath(filename)
	if err != nil {
		return nil, err
	}

	return os.Stat(fullpath)
}

func (fs *Local) openFile(filename string, flag int, perm os.FileMode) (extfs.File, error) {
	if flag&os.O_CREATE != 0 {
		if err := fs.createDir(filename); err != nil {
			return nil, err
		}
	}

	return os.OpenFile(filename, flag, perm)
}

func (fs *Local) createDir(fullpath string) error {
	dir := filepath.Dir(fullpath)
	if dir != "." {
		if err := os.MkdirAll(dir, defaultDirectoryMode); err != nil {
			return err
		}
	}

	return nil
}

func (fs *Local) underlyingPath(filename string) (string, error) {
	if isCrossBoundaries(filename) {
		return "", extfs.ErrCrossedBoundary
	}

	return filepath.Join(fs.base, filename), nil
}

func isCrossBoundaries(path string) bool {
	path = filepath.ToSlash(path)
	path = filepath.Clean(path)

	return strings.HasPrefix(path, ".."+string(filepath.Separator))
}
