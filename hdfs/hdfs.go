/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package hdfs

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/colinmarc/hdfs/v2"
	"github.com/colinmarc/hdfs/v2/hadoopconf"
	"github.com/rkcloudchain/extfs"
	"github.com/rkcloudchain/extfs/util"
)

const (
	defaultDirectoryMode = 0755
	defaultCreateMode    = 0666
)

// hadoop s a filesystem based on the hadoop filesystem.
type hadoop struct {
	client *hdfs.Client
	base   string
}

// New returns a hadoop filesystem.
func New(baseDir string, cfg *extfs.Config) (extfs.Filesystem, error) {
	hadoopCfg, err := hadoopconf.LoadFromEnvironment()
	if err != nil {
		return nil, err
	}

	options := hdfs.ClientOptionsFromConf(hadoopCfg)
	options.UseDatanodeHostname = cfg.UseDatanodeHostname
	if len(cfg.Addresses) != 0 {
		options.Addresses = append(options.Addresses, cfg.Addresses...)
	}
	if cfg.User != "" {
		options.User = cfg.User
	} else {
		u, err := user.Current()
		if err != nil {
			options.User = "root"
		} else {
			options.User = u.Username
		}
	}

	client, err := hdfs.NewClient(options)
	if err != nil {
		return nil, err
	}

	return &hadoop{client, baseDir}, nil
}

func (fs *hadoop) Create(filename string) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return fs.createFile(fullpath)
}

func (fs *hadoop) Open(filename string) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return fs.openFile(fullpath)
}

func (fs *hadoop) OpenFile(filename string, flag int, perm os.FileMode) (extfs.File, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	// github.com/colinmarc/hdfs is temporarily not supported
	if flag&os.O_TRUNC != 0 {
		return nil, errors.New("HDFS does not support truncate operation")
	}

	accMode := flag & syscall.O_ACCMODE
	if accMode == os.O_RDWR {
		return nil, errors.New("HDFS file can only be opened as read-only or write-only")
	}
	if accMode == os.O_WRONLY && flag&os.O_CREATE == 0 && flag&os.O_APPEND == 0 {
		return nil, errors.New("HDFS file can only be append written")
	}

	if accMode == os.O_RDONLY {
		return fs.openFile(fullpath)
	}

	if flag&os.O_CREATE != 0 {
		_, err := fs.client.Stat(fullpath)
		if err == nil && flag&os.O_EXCL != 0 {
			return nil, os.ErrExist
		}
		if os.IsNotExist(err) {
			return fs.createFile(fullpath)
		}
	}

	return fs.appendFile(fullpath)
}

func (fs *hadoop) Remove(filename string) error {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return err
	}

	return fs.client.Remove(fullpath)
}

func (fs *hadoop) RemoveAll(path string) error {
	fullpath, err := util.UnderlyingPath(fs.base, path)
	if err != nil {
		return err
	}

	return fs.client.Remove(fullpath)
}

func (fs *hadoop) Rename(oldpath, newpath string) error {
	var err error
	oldpath, err = util.UnderlyingPath(fs.base, oldpath)
	if err != nil {
		return err
	}

	newpath, err = util.UnderlyingPath(fs.base, newpath)
	if err != nil {
		return err
	}

	return fs.client.Rename(oldpath, newpath)
}

func (fs *hadoop) Stat(filename string) (os.FileInfo, error) {
	fullpath, err := util.UnderlyingPath(fs.base, filename)
	if err != nil {
		return nil, err
	}

	return fs.client.Stat(fullpath)
}

func (fs *hadoop) ReadDir(path string) ([]os.FileInfo, error) {
	fullpath, err := util.UnderlyingPath(fs.base, path)
	if err != nil {
		return nil, err
	}

	return fs.client.ReadDir(fullpath)
}

func (fs *hadoop) MkdirAll(path string, perm os.FileMode) error {
	fullpath, err := util.UnderlyingPath(fs.base, path)
	if err != nil {
		return err
	}

	return fs.client.MkdirAll(fullpath, defaultDirectoryMode)
}

func (fs *hadoop) Chmod(name string, mode os.FileMode) error {
	fullpath, err := util.UnderlyingPath(fs.base, name)
	if err != nil {
		return err
	}

	return fs.client.Chmod(fullpath, mode)
}

func (fs *hadoop) Chtimes(name string, atime time.Time, mtime time.Time) error {
	fullpath, err := util.UnderlyingPath(fs.base, name)
	if err != nil {
		return err
	}

	return fs.client.Chtimes(fullpath, atime, mtime)
}

func (fs *hadoop) Close() error {
	return fs.client.Close()
}

func (fs *hadoop) createFile(fullpath string) (extfs.File, error) {
	dir := filepath.Dir(fullpath)
	err := fs.client.MkdirAll(dir, defaultDirectoryMode)
	if err != nil {
		return nil, err
	}

	fw, err := fs.client.Create(fullpath)
	if err != nil {
		return nil, err
	}

	return newFile(nil, fw), nil
}

func (fs *hadoop) openFile(fullpath string) (extfs.File, error) {
	fr, err := fs.client.Open(fullpath)
	if err != nil {
		return nil, err
	}

	return newFile(fr, nil), nil
}

func (fs *hadoop) appendFile(fullpath string) (extfs.File, error) {
	fw, err := fs.client.Append(fullpath)
	if err != nil {
		return nil, err
	}

	return newFile(nil, fw), nil
}
