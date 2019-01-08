/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package local

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/rkcloudchain/extfs"
	"github.com/rkcloudchain/extfs/util"
)

const (
	defaultDirectoryMode = 0755
	defaultCreateMode    = 0666
)

// local is a filesystem based on the local filesystem.
type local struct {
	base string
}

// New returns a local filesystem.
func New(baseDir string) extfs.Filesystem {
	return &local{baseDir}
}

// Create ...
func (fs *local) Create(filename string) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return fs.openFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, defaultCreateMode)
}

// Open ...
func (fs *local) Open(filename string) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return os.Open(fullpath)
}

// OpenFile ...
func (fs *local) OpenFile(filename string, flag int, perm os.FileMode) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return fs.openFile(fullpath, flag, perm)
}

// Rename ...
func (fs *local) Rename(from, to string) error {
	var err error
	from, err = util.UnderlyingPath(fs.base, from)
	if err != nil {
		return err
	}

	to, err = util.UnderlyingPath(fs.base, to)
	if err != nil {
		return err
	}

	if err := fs.createDir(to); err != nil {
		return err
	}

	return os.Rename(from, to)
}

// Remove ...
func (fs *local) Remove(filename string) error {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return err
	}

	return os.Remove(fullpath)
}

// RemoveAll ...
func (fs *local) RemoveAll(path string) error {
	fullpath, err := util.UnderlyingPath(fs.base, path)
	if err != nil {
		return err
	}

	return os.RemoveAll(fullpath)
}

// ReadDir ...
func (fs *local) ReadDir(path string) ([]os.FileInfo, error) {
	fullpath, err := util.UnderlyingPath(fs.base, path)
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
func (fs *local) MkdirAll(path string, perm os.FileMode) error {
	fullpath, err := util.UnderlyingPath(fs.base, path)
	if err != nil {
		return err
	}

	return os.MkdirAll(fullpath, defaultDirectoryMode)
}

// Stat ...
func (fs *local) Stat(filename string) (os.FileInfo, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return os.Stat(fullpath)
}

func (fs *local) Chmod(name string, mode os.FileMode) error {
	fullpath, err := util.UnderlyingPath(fs.base, name)
	if err != nil {
		return err
	}

	return os.Chmod(fullpath, mode)
}

func (fs *local) Chtimes(name string, atime time.Time, mtime time.Time) error {
	fullpath, err := util.UnderlyingPath(fs.base, name)
	if err != nil {
		return err
	}

	return os.Chtimes(fullpath, atime, mtime)
}

// Close ...
func (fs *local) Close() error {
	return nil
}

func (fs *local) openFile(filename string, flag int, perm os.FileMode) (extfs.File, error) {
	if flag&os.O_CREATE != 0 {
		if err := fs.createDir(filename); err != nil {
			return nil, err
		}
	}

	return os.OpenFile(filename, flag, perm)
}

func (fs *local) createDir(fullpath string) error {
	dir := filepath.Dir(fullpath)
	if dir != "." {
		if err := os.MkdirAll(dir, defaultDirectoryMode); err != nil {
			return err
		}
	}

	return nil
}
